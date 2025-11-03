package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "bbbab"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "bbaaba"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "aaaa"
	expect := -1
	runSample(t, s, expect)
}