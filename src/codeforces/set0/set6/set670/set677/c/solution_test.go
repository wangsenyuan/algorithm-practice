package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "z"
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "V_V"
	expect := 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "Codeforces"
	expect := 130653412
	runSample(t, s, expect)
}