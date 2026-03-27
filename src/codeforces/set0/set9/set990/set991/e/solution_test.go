package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2028, 13)
}

func TestSample2(t *testing.T) {
	runSample(t, 79, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000000000000, 18)
}
