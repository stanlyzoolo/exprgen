package exprgen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Generate provides the creation of a random expression according
// to the input length parameter. 
func Generate(r uint8) string {

	rand.Seed(time.Now().UnixNano())

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

	operand := func() {
		fmt.Fprint(&b, space(), any("0123456789"), space())
	}

	operator := func() {
		fmt.Fprint(&b, space(), any("+-"), space())
	}

	operand() // init builder by first digit

	for i := 0; uint8(i) < r; i++ {
		operator()
		operand()
	}

	return b.String()

}
