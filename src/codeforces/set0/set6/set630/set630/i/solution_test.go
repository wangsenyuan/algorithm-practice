package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 24)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 132)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 672)
}
