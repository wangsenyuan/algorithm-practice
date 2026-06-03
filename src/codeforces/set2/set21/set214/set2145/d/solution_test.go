package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect bool) {
	res := solve(n, k)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}

	if !expect {
		return
	}

	var arr []int
	for i := 0; i+1 < n; i++ {
		if res[i] > res[i+1] {
			arr = append(arr, i)
		}
	}

	sorted := slices.Clone(res)

	slices.Sort(sorted)
	for i := range n {
		if sorted[i] != i+1 {
			t.Fatalf("Sample result %v, is not a permutation", res)
		}
	}

	var sum int
	for i, v := range arr {
		u := n - v - 1
		if i+1 < len(arr) {
			u = arr[i+1] - v
		}
		sum += (v + 1) * u
	}

	if sum != k {
		t.Fatalf("Sample result %v, is not valid", res)
	}
}

func TestSample1(t *testing.T) {
	// 3 1 4 2
	// arr = []int{1, 3}
	// 1 * 2 + 3 * 1 = 5
	runSample(t, 4, 5, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 10, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 8, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 3, 1, false)
}


func TestSample5(t *testing.T) {
	runSample(t, 2, 1, true)
}

