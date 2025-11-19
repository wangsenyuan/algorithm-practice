package main

import "testing"

func runSample(t *testing.T, n int, C int, expect int) {
	res := solve(n, C)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 5
	C := 1
	expect := 5
	runSample(t, n, C, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	C := 2
	expect := 5
	runSample(t, n, C, expect)
}

func TestSample3(t *testing.T) {
	n := 11
	C := 5
	expect := 4367
	runSample(t, n, C, expect)
}

func TestSample4(t *testing.T) {
	n := 37
	C := 63
	expect := 230574
	runSample(t, n, C, expect)
}
