package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)
	if res != expect {
		t.Fatalf("expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 1, 3)
}
