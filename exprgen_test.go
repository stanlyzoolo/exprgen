package exprgen

import (
	. "github.com/stanlyzoolo/basiccalc"
	"testing"
	"testing/quick"
)

func countDigits(s string) uint8 {
	digits := "0123456789"
	var b uint8 = 0

	for _, r := range s {
		for _, v := range digits {
			if r == v {
				b++
			}
		}
	}
	return b
}

func countOperands(s string) uint8 {
	ops := "+-"

	var a uint8 = 0
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

	genOp := func(r uint8) bool {
		t.Logf("genOp parameter: %v", r)
		return countOperands(Generate(r)) == r
	}
	errOp := quick.Check(genOp, nil)

	if errOp != nil {
		t.Errorf("failed Generate(); %s", errOp)
	}

	genDig := func(r uint8) bool {
		t.Logf("genDig parameter: %v", r)

		return countDigits(Generate(r)) == r+1
	}
	errDig := quick.Check(genDig, nil)

	if errDig != nil {
		t.Errorf("failed Generate(); %s", errDig)
	}

	var r uint8 = 2
	_, err := Eval(Generate(r))
	if err != nil {
		t.Errorf("failed Eval(); %s", err)
	}

}
