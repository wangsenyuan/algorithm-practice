package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, s int, E [][]int, Q [][]int, expect []int) {
	res := solve(n, s, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, s := 5, 4
	E := [][]int{
		{3, 1},
		{2, 4},
		{5, 1},
		{1, 4},
	}
	Q := [][]int{
		{2, 4},
		{3, 2},
		{3, 5},
		{1, 2},
		{4, 5},
	}
	expect := []int{4, 4, 1, 4, 4}
	runSample(t, n, s, E, Q, expect)
}
