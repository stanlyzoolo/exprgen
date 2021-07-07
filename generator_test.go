package exprgen

import (
	"math/rand"
	"testing"
	"time"
)

func TestSymbolGen(t *testing.T) {
	var testSymbols []string

	testSymbols = []string{"2", "4", "+"}

	var i, v, op, sd interface{}
	i = "-"
	op = operators[1]

	v = "5"
	sd = singledigits[5]

	if i.(string) != op.(string) {
		t.Error("wrong type in operators value")
	}

	if v.(string) != sd.(string) {
		t.Error("wrong type in singledigits value")
	}

	s, err := symbolGen()
	if len(s) < len(testSymbols) && err != nil {
		t.Errorf("failed randSymbols(); %s", err)
	}

}

func TestInputMaker(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var testSymbols []string

	testSymbols = append(testSymbols,
		singledigits[rand.Intn(len(singledigits))],
		operators[rand.Intn(len(operators))],
		singledigits[rand.Intn(len(singledigits))],
	)

	InputMaker(testSymbols)

}
