package main

import "testing"

func runSample(t *testing.T, k int, a int, b int, expect int) {
	res := solve(k, a, b)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	k := 1
	a := 1
	b := 10
	expect := 10
	runSample(t, k, a, b, expect)
}

func TestSample2(t *testing.T) {
	k := 2
	a := -4
	b := 4
	expect := 5
	runSample(t, k, a, b, expect)
}
