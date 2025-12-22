package main

import "testing"

func TestSample1(t *testing.T) {
	cups := [][]int{
		{4, 10},
		{6, 8},
		{3, 5},
		{4, 14},
		{10, 9},
		{4, 20},
	}

	solver := NewSolver(cups)

	queries := [][]int{
		{1, 25},
		{6, 30},
		{5, 8},
		{3, 13},
		{2, 8},
	}
	expect := []int{5, 0, 5, 4, 2}

	for i, cur := range queries {
		res := solver.Query(cur[0], cur[1])
		if res != expect[i] {
			t.Fatalf("Sample expect %d, but got %d", expect[i], res)
		}
	}
}
