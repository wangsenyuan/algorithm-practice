package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, m int, n int, expect float64) {
	res := solve(m, n)
	if math.Abs(res-expect) > 1e-6 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 1, 3.500000000000)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 3, 4.958333333333)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 2, 1.750000000000)
}