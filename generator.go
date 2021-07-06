package main

import (
	"fmt"
	"math/rand"
	"github.com/stanlyzoolo/basiccalc"
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

func randRune(s map[rune]int) rune {
	var seq []rune
	for k := range s {
		seq = append(seq, k)
	}
	return seq[rand.Intn(len(seq))]
}

func setArgument() rune {
	return randRune(singledigits)
}

func setOperator() rune {
	return randRune(operators)
}

func makeExpression() string {
	var sequence []rune
	sequence = append(sequence, setArgument(), setOperator())

	if len(sequence) < 3 {
		sequence = append(sequence, setArgument())
	}

	return string(sequence)
}

func main() {


	srt := makeExpression()
	fmt.Println(srt)
	fmt.Println(basiccalc.Eval(srt))
}
