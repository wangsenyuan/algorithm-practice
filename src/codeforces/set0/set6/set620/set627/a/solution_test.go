package main

import "testing"

func runSample(t *testing.T, s int, x int, expect int) {
	res := solve(s, x)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 9, 5, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 3, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 2, 0)
}
