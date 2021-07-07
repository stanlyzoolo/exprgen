package exprgen

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// operators map represents mathematical operators.
var operators = map[int]string{
	0: "+",
	1: "-",
}

// singledigits map represent single digits from 0 to 9.
var singledigits = map[int]string{
	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5",
	6: "6", 7: "7", 8: "8", 9: "9",
}

// symbolGen generates a random arguments and operator and
// return symbols, as a slice of strings.
func symbolGen() ([]string, error) {
	// symbols represent a storage for random strings.
	var symbols []string

	rand.Seed(time.Now().UnixNano())

	symbols = append(symbols,
		singledigits[rand.Intn(len(singledigits))],
		operators[rand.Intn(len(operators))],
		singledigits[rand.Intn(len(singledigits))],
	)

	return symbols, errors.New("failed randSymbols()")

}

// InputMaker build a string from symbols, that represent an expression.
func InputMaker(symbols []string) string {
	var b strings.Builder
	for _, s := range symbols {
		b.WriteString(s)
	}

	return b.String()
}
