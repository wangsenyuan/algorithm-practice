package main

import "testing"

func runSample(t *testing.T, n int, x int, k int, expect int) {
	res := solve(n, x, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 0, 1, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 4, 25, 25)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 7, 2, 582)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000000000, 40, 1000000000, 514035484)
}