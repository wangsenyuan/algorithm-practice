package main

import "testing"

func runSample(t *testing.T, n int, l int, r int, expect int) {
	res := solve(n, l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 3, 10, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 2, 5, 4)
}
