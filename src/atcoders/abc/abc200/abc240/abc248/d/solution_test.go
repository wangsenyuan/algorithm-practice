package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, Q [][]int, expect []int) {
	res := solve(A, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 1, 4, 1, 5}
	Q := [][]int{
		{1, 5, 1},
		{2, 4, 3},
		{1, 5, 2},
		{1, 3, 3},
	}
	expect := []int{2, 0, 0, 1}
	runSample(t, A, Q, expect)
}
