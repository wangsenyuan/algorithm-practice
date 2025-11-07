package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "3+5*7+8*4", 303)
}

func TestSample2(t *testing.T) {
	runSample(t, "2+3*5", 25)
}

func TestSample3(t *testing.T) {
	runSample(t, "3*4*5", 60)
}
