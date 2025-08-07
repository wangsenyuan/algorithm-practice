package main

import "testing"

func runSample(t *testing.T, x int, y int, m int, expect int) {
	res := solve(x, y, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 5, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, -1, 4, 15, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 0, -1, 5, -1)
}
