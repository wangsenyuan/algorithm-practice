package main

import (
	"slices"
	"testing"
)

func runSample(t *testing.T, n int, k int, expect []int) {
	ans := solve(n, k)
	if !slices.Equal(ans, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	n, k := 4, 3
	expect := []int{1, 3, 2, 4}
	runSample(t, n, k, expect)
}

func TestSample2(t *testing.T) {
	n, k := 10, 1
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	runSample(t, n, k, expect)
}
