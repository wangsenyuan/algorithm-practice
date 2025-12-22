package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, E [][]int, Q []int, expect []bool) {
	res := solve(n, E, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	E := [][]int{
		{1, 2, 2},
	}
	Q := []int{2}
	expect := []bool{true}
	runSample(t, n, E, Q, expect)
}
