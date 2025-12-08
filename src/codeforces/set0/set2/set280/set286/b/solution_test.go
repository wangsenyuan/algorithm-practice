package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 2
	expect := []int{2, 1}
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 3
	expect := []int{1, 3, 2}
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 4
	expect := []int{4, 2, 3, 1}
	runSample(t, n, expect)
}

func TestSample4(t *testing.T) {
	n := 10
	expect := []int{10, 2, 9, 6, 7, 1, 8, 3, 5, 4}
	runSample(t, n, expect)
}