package main

import "testing"

func runSample(t *testing.T, a int, m int, expect int) {
	res := solve(a, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 10, 7)
}
