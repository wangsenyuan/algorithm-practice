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
	s := `2 2
..
.X
`
	runSample(t, s, 0.888888888889)
}

func TestSample2(t *testing.T) {
	s := `3 3
...
.X.
...
`
	runSample(t, s, 2.0)
}
