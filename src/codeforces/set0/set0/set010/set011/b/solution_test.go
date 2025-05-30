package main

import "testing"

func runSample(t *testing.T, x int, expect int) {
	res := solve(x)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, -1000000000, 44723)
}
