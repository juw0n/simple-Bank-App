package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate a random integer btw min and max value
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// generate a random owner name
func RandomOwner() string {
	return RandomString(5)
}

// generate a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 5000)
}

// renerate a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "NGN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
