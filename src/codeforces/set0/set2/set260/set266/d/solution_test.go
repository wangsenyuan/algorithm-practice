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
	if math.Abs(res-expect) > 1e-9 {
		t.Errorf("Sample expect %.10f, but got %.10f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 1
1 2 1
`
	expect := 0.5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 2 1
2 3 1
1 3 1
`
	expect := 1.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 2
1 2 100
2 3 1
`
	expect := 50.50
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 4
3 4 10
2 4 6
2 3 9
4 1 7
`
	expect := 8.50
	runSample(t, s, expect)
}

func TestCounterExample(t *testing.T) {
	s := `4 4
2 1 4
3 2 9
4 2 7
3 4 6
`
	expect := 8.0
	runSample(t, s, expect)
}
