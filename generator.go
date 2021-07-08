package exprgen

import (
	"math/rand"
	"strings"
	"time"
)

func Generate(r uint) string {

	ops := "+-"
	digits := "0123456789"

	rand.Seed(time.Now().UnixNano()) //

	any := func(data string) string {
		return string(data[rand.Intn(len(data))])
	}

	space := func() string {
		if rand.Intn(2) == 0 {
			return ""
		}
		return " "
	}

	var b strings.Builder

	decorator := func(s string) {
		b.WriteString(space())
		b.WriteString(s)
		b.WriteString(space())
	}

	operand := func() {
		decorator(any(digits))
	}
	operator := func() {
		decorator(any(ops))
	}

	operand() // init builder by first digit

	for i := 0; uint(i) < r; i++ {
		operator()
		operand()
	}

	return b.String()
}
