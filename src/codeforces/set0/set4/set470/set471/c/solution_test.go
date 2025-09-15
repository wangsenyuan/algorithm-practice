package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 26, 2)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000000000000, 272165)
}
