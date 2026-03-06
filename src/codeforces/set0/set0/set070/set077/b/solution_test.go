package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, a int, b int, expect float64) {
	res := solve(a, b)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 4, 2
	expect := 0.625
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b := 1, 2
	expect := 0.5312500000
	runSample(t, a, b, expect)
}
