package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, w int, b int, expect float64) {
	ans := solve(w, b)

	if math.Abs(ans-expect) > 1e-9 {
		t.Fatalf("Sample expect %.10f, but got %.10f", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 3, 0.5)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 5, 0.658730159)
}
