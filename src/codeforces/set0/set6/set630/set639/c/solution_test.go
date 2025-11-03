package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 1000000000
10 -9 -3 5
`

	// P(2) = 10 - 18 - 12 + 40 = 20
	// Q(2) = 只要有个系数能变成20就可以了
	//       -10， -19, -8
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 12
10 -9 -3 5
`

	expect := 2
	runSample(t, s, expect)
}
