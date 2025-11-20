package main

import "testing"

func runSample(t *testing.T, m int, expect int) {
	res := solve(m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 54)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, -1)
}

