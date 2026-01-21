package main

import (
	"bufio"
	"math/big"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	w := big.NewInt(int64(expect))

	if w.Cmp(&res) != 0 {
		t.Errorf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2 3 1 3
`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
6 6 4 4
`
	expect := -8
	runSample(t, s, expect)
}
