package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, n int, k int, E [][]int, expect []int) {
	res := solve(n, k, E)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 3, 2
	E := [][]int{
		{1, 3, 1},
	}
	expect := []int{0, 1, 1}
	runSample(t, n, k, E, expect)
}

func TestSample2(t *testing.T) {
	n, k := 4, 1
	E := [][]int{
		{1, 2, 3},
		{2, 4, 5},
		{3, 4, 7},
	}
	expect := []int{0, 1, 4, 6}
	runSample(t, n, k, E, expect)
}

func TestSample3(t *testing.T) {
	n, k := 5, 2
	E := [][]int{
		{2, 1, 33},
		{1, 5, 93},
		{5, 3, 48},
		{2, 3, 21},
		{4, 2, 1},
	}
	expect := []int{0, 1, 2, 2, 3}
	runSample(t, n, k, E, expect)
}
