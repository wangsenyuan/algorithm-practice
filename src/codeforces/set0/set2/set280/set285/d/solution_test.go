package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 15, 150347555)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 18)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 1800)
}
