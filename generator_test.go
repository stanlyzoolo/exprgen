package exprgen

import (
	"testing"
	"testing/quick"

	. "github.com/stanlyzoolo/basiccalc"
)


func TestGenerate(t *testing.T) {
	var r uint = 1000 // any uint value

	gen := func(r uint) bool {
		r = 100
		return uint(len(Generate(r))) == r
	}

	var i int = 10 // any int value
	err := quick.Check(gen, &quick.Config{MaxCount: i})

	if err != nil {
		t.Errorf("failed Generate(); %s", err)
	}

	_, err = Eval(Generate(r))
	if err != nil {
		t.Errorf("failed Eval(); %s", err)
	}

}

