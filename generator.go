package exprgen

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// operators map represents mathematical operators.
var operators = map[int]string{
	0: "+",
	1: "-",
}

// singledigits map represent single digits from 0 to 9.
var singledigits = map[int]string{
	0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5",
	6: "6", 7: "7", 8: "8", 9: "9",
}

func randLenExpression(odd int) (string, error) {
	rand.Seed(time.Now().UnixNano())

	if odd%2 == 0 {
		return "", errors.New("invalid input argument; enter odd digit")
	}

	var symbols []string

	for i := odd; i > 0; i-- {
		symbols = append(symbols,
			singledigits[rand.Intn(len(singledigits))],
			operators[rand.Intn(len(operators))],
		)

	}

	symbols = symbols[:odd]

	var b strings.Builder
	for _, v := range symbols {
		b.WriteString(v)
	}

	return b.String(), nil

}
