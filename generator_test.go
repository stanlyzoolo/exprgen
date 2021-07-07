package exprgen

import (
	// "math/rand"
	"testing"
	// "time"
)

func TestRandLenExpression(t *testing.T) {
	
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

	evenDigit := 8

	v, err := randLenExpression(evenDigit)

	if err == nil {
		t.Errorf("failed randLenExpression() with %v input argument; %s", v, err)

	}

	oddDigit := 7
	v, err = randLenExpression(oddDigit)

	if err != nil {
		t.Errorf("failed randLenExpression() with %v input argument; %s", v, err)

	}


	

}
