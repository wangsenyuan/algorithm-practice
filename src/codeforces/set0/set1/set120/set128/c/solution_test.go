package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 4, 1, 9)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 7, 2, 75)
}

func TestSample4(t *testing.T) {
	runSample(t, 999, 999, 499, 1)
}
