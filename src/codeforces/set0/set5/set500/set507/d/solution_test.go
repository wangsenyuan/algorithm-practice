package main

import "testing"

func runSample(t *testing.T, n int, k int, m int, expect int) {
	res := solve(n, k, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 2, 1000, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2, 1000, 45)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 3, 1103, 590)
}

func TestSample4(t *testing.T) {
	runSample(t, 183 ,3, 46847167, 29891566)
}
