package main

import "testing"

func runSample(t *testing.T, s string, expect bool) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ababa"
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "baaba"
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "?b?ab"
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "aa?b?b"
	expect := false
	runSample(t, s, expect)
}
