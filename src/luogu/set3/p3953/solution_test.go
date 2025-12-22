package main

import "testing"

func runSample(t *testing.T, n int, k int, p int, edges [][]int, expect int) {
	res := solve(n, k, p, edges)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k, p := 5, 2, 10
	edges := [][]int{
		{1, 2, 1},
		{2, 4, 0},
		{4, 5, 2},
		{2, 3, 2},
		{3, 4, 1},
		{3, 5, 2},
		{1, 5, 3},
	}
	expect := 3
	runSample(t, n, k, p, edges, expect)
}

func TestSample2(t *testing.T) {
	n, k, p := 2, 0, 10
	edges := [][]int{
		{1, 2, 0},
		{2, 1, 0},
	}
	expect := -1
	runSample(t, n, k, p, edges, expect)
}
