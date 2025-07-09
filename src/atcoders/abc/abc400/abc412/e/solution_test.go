package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 12, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 123456789, 123456789, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 99999990000000, 100000000000000, 310458)
}
