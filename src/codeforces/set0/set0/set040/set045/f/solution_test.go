package main

import "testing"

func runSample(t *testing.T, m int, n int, expect int) {
	res := solve(m, n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// (3, 3) (0, 0)
	// (1, 3) (2, 0)
	// (2, 3),(1, 0)
	// 
	runSample(t, 3, 2, 11)
}
