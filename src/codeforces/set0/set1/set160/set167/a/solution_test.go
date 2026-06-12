package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []float64) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	for i, x := range expect {
		y := res[i]
		if math.Abs(x-y)/max(1, x) > 1e-4 {
			t.Fatalf("Sample result expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 10 10000
0 10
5 11
1000 1
`
	expect := []float64{
		1000.5000000000,
		1000.5000000000,
		11000.0500000000,
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2 26
28 29
`
	expect := []float64{
		33.0990195136,
	}
	runSample(t, s, expect)
}
