package main

import "testing"

func runSample(t *testing.T, a, b int, expect int) {
	res := solve(a, b)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 1, 2)
}

func TestSample2(t *testing.T) {
	// 4 * 5 * 6
	// 2 * 2 * 5 * 2 * 3
	runSample(t, 6, 3, 5)
}
