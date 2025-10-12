package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "11011"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "10000011"
	expect := 5
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "11"
	expect := 2
	runSample(t, s, expect)
}
