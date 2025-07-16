package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "chipsy48.32televizor12.390", "12.438.32")
}

func TestSample2(t *testing.T) {
	runSample(t, "a1b2c3.38", "6.38")
}

func TestSample3(t *testing.T) {
	runSample(t, "aa0.01t0.03", "0.04")
}

func TestSample4(t *testing.T) {
	runSample(t, "test0.50test0.50", "1")
}
