package main

import "testing"

func runSample(t *testing.T, l, r int, expect int) {
	res := solve(l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 16, 31)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1, 0)
}