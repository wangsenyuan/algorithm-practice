package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	_, x, y, best, res := drive(reader)

	if best != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, best)
	}

	n := len(x)

	var sum int
	for i := range n {
		z := res[id(x[i])]
		if z == y[i] {
			sum++
		}
	}

	if sum != best {
		t.Fatalf("Sample expect %d, but got %d", expect, sum)
	}
}

func TestSample1(t *testing.T) {
	s := `10 2
aaabbbaaab
bbbbabbbbb
`
	expect := 7
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 2
aaaaaaabbb
bbbbaaabbb
`
	expect := 6
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9 4
dacbdacbd
acbdacbda
`
	expect := 9
	runSample(t, s, expect)
}
