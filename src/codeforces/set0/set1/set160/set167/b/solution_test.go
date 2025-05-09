package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	if math.Abs(res-expect) > 1e-7 {
		t.Errorf("Sample expect %.7f, but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 1 0
10 20 30
-1 -1 2
`
	expect := 0.3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 1 1
100
123
`
	expect := 1.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 0 0
7
-1
`
	expect := 0.93
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `9 9 2
91 96 99 60 42 67 46 39 62
5 -1 2 -1 -1 -1 7 -1 3
`
	expect := 0.016241917181
	runSample(t, s, expect)
}
