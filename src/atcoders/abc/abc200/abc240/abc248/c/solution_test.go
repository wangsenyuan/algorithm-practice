package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, expect int) {
	res := solve(n, m, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 4, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 31, 41, 592, 798416518)
}
