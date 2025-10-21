package main

import "testing"

func runSample(t *testing.T, r int, g int, b int, expect int) {
	res := solve(r, g, b)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 4, 3, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 3, 3, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 0, 1, 1000000000, 1)
}
