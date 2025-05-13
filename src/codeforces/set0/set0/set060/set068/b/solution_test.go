package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if math.Abs(res-expect) > 1e-8 {
		t.Errorf("Sample expect %.8f, but got %.8f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 50
4 2 1
`
	expect := 2.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 90
1 11
`
	expect := 1.909090909
	runSample(t, s, expect)
}
