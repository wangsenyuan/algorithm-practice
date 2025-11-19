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
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 3
`
	expect := 1.0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
1 2 3 4
`
	expect := 2.0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
1 10 2 9 3 8 4 7 5 6
`
	expect := 4.5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
0 0 0 0 0 0 0 0 0 0
	`
	expect := 0.0
	runSample(t, s, expect)
}
