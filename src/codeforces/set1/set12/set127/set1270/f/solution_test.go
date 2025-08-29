package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	ans := solve(s)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "111", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "1111100000", 25)
}
