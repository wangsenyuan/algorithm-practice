package main

import "testing"

func runSample(t *testing.T, p int, k int, expect int) {
	res := solve(p, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10007, 25, 100140049)
}

func TestSample2(t *testing.T) {
	runSample(t, 65213, 29960, 65213)
}
