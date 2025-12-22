package main

import (
	"math"
	"testing"
)

func runSample(t *testing.T, n int, edges [][]int, h int, expect float64) {
	res := solve(n, edges, h)

	if math.Abs(res-expect) > 1e-5 {
		t.Fatalf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, h := 5, 100
	edges := [][]int{
		{1, 2, 20, 5},
		{1, 3, 20, 5},
		{1, 4, 20, 5},
		{1, 5, 20, 5},
		{2, 3, 23, 1},
	}
	var expect float64 = 1.0625
	runSample(t, n, edges, h, expect)
}
