package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, a int, b int, expect float64) {
	res := solve(a, b)
	if math.Abs(res-expect)/max(1, expect) > 1e-6 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a, b := 3, 1
	expect := 1.0
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a, b := 1, 3
	expect := -1.0
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a, b := 4, 1
	expect := 1.25
	runSample(t, a, b, expect)
}
