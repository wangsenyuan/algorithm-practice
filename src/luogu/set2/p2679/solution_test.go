package main

import "testing"

func runSample(t *testing.T, a string, b string, k int, expect int) {
	res := solve(a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "aabaab"
	b := "aab"
	k := 1
	expect := 2
	runSample(t, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	a := "aabaab"
	b := "aab"
	k := 1
	expect := 2
	runSample(t, a, b, k, expect)
}
