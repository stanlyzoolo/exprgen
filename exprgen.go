package exprgen

import (
	"math/rand"
	"strings"
	// "github.com/stanlyzoolo/basiccalc"
)

// operators is a map where keys represent mathematical operators as a string
//  type and values represent the corresponding function.
var operators = map[rune]int{
	'+': '+',
	'-': '-',
}

// singledigits is a map where keys represent single digits
//  as a string type and values represent them in type int.
var singledigits = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

func randRune(s map[rune]int) string {
	var seq []string
	for k := range s {
		seq = append(seq, string(k))
	}
	return seq[rand.Intn(len(seq))]
}

func setArgument() string {
	return randRune(singledigits)
}

func setOperator() string {
	return randRune(operators)
}

func Generate(r uint) []string {

	var b strings.Builder
	b.WriteString(setArgument())
	b.WriteString(setOperator())
	b.WriteString(setArgument())

	var sequence []string
	sequence = append(sequence, b.String())

	// if len(sequence) < 3 {
	// 	sequence = append(sequence, setArgument())
	// }

	return sequence
}
