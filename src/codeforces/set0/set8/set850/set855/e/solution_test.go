package main

import "testing"

func runSample(t *testing.T, b int, l int, r int, expect int) {
	res := solve(b, l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 4, 9, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 1, 10, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 1, 100, 21)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 1, 100, 4)
}