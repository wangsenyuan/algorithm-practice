package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{1, 2, 3}
	expect := 2
	runSample(t, a, expect)
}
