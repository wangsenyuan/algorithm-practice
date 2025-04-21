package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 14, 15}
	expect := 2044
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1001, 5, 1000000, 1000000000, 100000}
	expect := 625549048
	runSample(t, a, expect)
}
