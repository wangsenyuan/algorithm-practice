package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 2, 5)
}

func TestSample3(t *testing.T) {
	// 0 : 1
	runSample(t, 1, 2, 1)
}
