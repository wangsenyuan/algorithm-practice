package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.Abs(res-expect)/max(1, expect) > 1e-9 {
		t.Fatalf("Sample expect %.10f, but got %.10f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
2 1 6
0 4 1
2 -1 3
1 -2 1
4 -1 1
`
	expect := 138.23007676
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
0 0 1
0 0 2
0 0 3
0 0 4
0 0 5
0 0 6
0 0 7
0 0 8

`
	expect := 289.02652413
	runSample(t, s, expect)
}
