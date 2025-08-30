package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "agerald.agapov1991@gmail.com", 18)
}

func TestSample2(t *testing.T) {
	runSample(t, "x@x.x@x.x_e_@r1.com", 8)
}

func TestSample3(t *testing.T) {
	runSample(t, "a___@1.r", 1)
}

func TestSample4(t *testing.T) {
	runSample(t, ".asd123__..@", 0)
}
