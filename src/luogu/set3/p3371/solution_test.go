package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, swaps [][]int, expect []int) {
	res := solve(n, swaps)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	swaps := [][]int{
		{1, 3},
		{1, 3},
		{1, 4},
	}
	expect := []int{4, 3, 2, 1, 5}
	runSample(t, n, swaps, expect)
}
