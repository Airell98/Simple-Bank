package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var TestDB *sql.DB
var mainError error

const (
	dbDriver = "postgres"
	dbSource = "postgresql://sofyan:postgres@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	TestDB, mainError = sql.Open(dbDriver, dbSource)

	if mainError != nil {
		log.Fatalln("Cannot connect to database:", mainError)
	}

	testQueries = New(TestDB)
	os.Exit(m.Run())
}
