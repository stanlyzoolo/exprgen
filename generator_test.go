package exprgen

import (
	"testing"
	"testing/quick"

	. "github.com/stanlyzoolo/basiccalc"
)

var r int = 1000 // any int value

func TestGenerate(t *testing.T) {

	gen := func(r int) bool {
		return len(Generate(r)) == r
	}

	err := quick.Check(gen, &quick.Config{MaxCount: r})

	if err != nil {
		t.Errorf("failed Generate(); %s", err)
	}

	_, err = Eval(Generate(r))
	if err != nil {
		t.Errorf("failed Eval(); %s", err)
	}

}
