package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if len(res) != expect {
		t.Fatalf("Sample expect %d, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 6
	expect := 4
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 10
	expect := 5
	runSample(t, n, expect)
}
