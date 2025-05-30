package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, A int, V int, L int, D int, W int, expect float64) {
	res := solve(A, V, L, D, W)
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 2, 1, 3, 2.5)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 70, 200, 170, 40, 8.965874696353)
}
