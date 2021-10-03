package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Returns random integer values
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Returns random strings
func GenerateString(n int) string {
	var sb strings.Builder
	l := len(alphabet)

	for i := 0; i < n; i++ {
		randomNum := rand.Intn(l)
		b := alphabet[randomNum]
		sb.WriteByte(b)
	}

	return sb.String()
}

// Returns random owner name from GenerateString function
func RandomOwner() string {
	return GenerateString(6)
}

// Returns random integer values from RandomInt function
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Returns random money currency
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "IDR", "RM", "CAD"}
	l := len(currencies)
	return currencies[rand.Intn(l)]
}
