package exprgen

import (
	. "github.com/stanlyzoolo/basiccalc"
	"testing"
	"testing/quick"
)

func countDigits(s string) int {
	digits := "0123456789"
	var b int = 0

	for _, r := range s {
		for _, v := range digits {
			if r == v {
				b++
			}
		}
	}
	return b
}

func countOperands(s string) int {
	ops := "+-"

	var a int = 0
	for _, r := range s {
		for _, v := range ops {
			if r == v {
				a++
			}
		}
	}
	return a
}

func TestGenerate(t *testing.T) {
	
	// var i int = 2 // any int value for &quick.Config{}

	genOp := func(r uint) bool {
		return countOperands(Generate(r)) == int(r+1)
	}
	errOp := quick.Check(genOp, nil) // &quick.Config{MaxCount: i}

	if errOp != nil {
		t.Errorf("failed Generate(); %s", errOp)
	}

	genDig := func(r uint) bool {
		return countDigits(Generate(r)) == int(r+2)
	}
	errDig := quick.Check(genDig, nil) // &quick.Config{MaxCount: i}

	if errDig != nil {
		t.Errorf("failed Generate(); %s", errDig)
	}

	var r uint = 2 // any int value
	_, err := Eval(Generate(r))
	if err != nil {
		t.Errorf("failed Eval(); %s", err)
	}

}
