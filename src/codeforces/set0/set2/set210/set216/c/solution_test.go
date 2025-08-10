package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)
	if len(res) != expect {
		t.Errorf("Sample expect %d, but got %d", expect, len(res))
	}
}

func TestSample1(t *testing.T) {
	// 1 1 4 5
	runSample(t, 4, 3, 2, 4)
}
