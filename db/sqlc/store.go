package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store struct provides all functions to execute all queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func CatchErr() {
	if err := recover(); err != nil {
		fmt.Println("Error occured", err.(error).Error())
	} else {
		fmt.Println("Application running prefectly")
	}
}

// execTx executes a function  within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

	// fmt.Printf("Starting transaction in execTx method: %+v \n", store.db)
	tx, err := store.db.BeginTx(ctx, nil)
	// fmt.Println("Starting Transaction error", err)

	if err != nil {
		// fmt.Println("error occured while starting transaction")
		return err
	}

	q := New(tx)

	errTx := fn(q)

	if errTx != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rb Error: %v, tx: %v", rbErr, errTx)
		}

		return errTx
	}

	return tx.Commit()
}

var txKey = struct{}{}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"ammount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

//TransferTx performs a money transfer from one account to another.
//It creates a transfer record, add account entries, and update account's balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	// fmt.Printf("argument parameter in TransferTx method: %+v \n", arg)
	// defer CatchErr()
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)

		fmt.Println(txName, "create transfer")
		result.Transfer, err = store.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})

		if err != nil {
			fmt.Println("error occured while invoking store.CreateTransfer method")
			return err
		}

		fmt.Println(txName, "create entry 1")
		result.FromEntry, err = store.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 2")
		result.ToEntry, err = store.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})

		fmt.Println(txName, "get account 1 for update")
		account1, err := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 1")
		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: account1.Balance - arg.Amount,
		})

		if err != nil {
			return err
		}

		fmt.Println(txName, "get account 2 for update")
		account2, err := q.GetAccountForUpdate(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 2")
		result.ToAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.ToAccountID,
			Balance: account2.Balance + arg.Amount,
		})

		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
