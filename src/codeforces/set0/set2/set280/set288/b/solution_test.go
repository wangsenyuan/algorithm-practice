package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2, 54)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 4, 1728)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 5, 16875)
}
