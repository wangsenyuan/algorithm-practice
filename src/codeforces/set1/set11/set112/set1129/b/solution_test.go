package main

import (
	"math/rand"
	"slices"
	"testing"
)

func runSample(t *testing.T, k int) {
	res := solve(k)
	if len(res) > 2*N {
		t.Fatalf("Sample result is too long, having %d", len(res))
	}
	if slices.Max(res) > X || slices.Min(res) < -X {
		t.Fatalf("Sample result is out of range, max = %d, min = %d", slices.Max(res), slices.Min(res))
	}
	f := func(arr []int) int {
		var res int
		var sum int
		k := -1
		for i, v := range arr {
			sum += v
			if sum < 0 {
				sum = 0
				k = i
			}
			res = max(res, (i-k)*sum)
		}
		return res
	}
	g := func(arr []int) int {
		var sum int
		for _, v := range arr {
			sum += v
		}
		return sum * len(arr)
	}

	if g(res)-f(res) != k {
		t.Fatalf("Sample %d result is not correct %v, f = %d, g = %d", k, res, f(res), g(res))
	}
}

func TestSample1(t *testing.T) {
	for range 10 {
		n := rand.Intn(1000) + 1
		runSample(t, n)
	}
}

func TestSample2(t *testing.T) {
	// X, 0......0, -1, 
	runSample(t, 1000000000)
}
