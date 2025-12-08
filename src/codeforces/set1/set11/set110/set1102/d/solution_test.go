package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "121", "021")
}

func TestSample2(t *testing.T) {
	runSample(t, "000000", "001122")
}

func TestSample3(t *testing.T) {
	runSample(t, "211200", "211200")
}

func TestSample4(t *testing.T) {
	runSample(t, "120110", "120120")
}
