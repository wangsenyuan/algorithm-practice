package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "MFM", 1)
}

func TestSample2(t *testing.T) {
	runSample(t, "FFMMM", 0)
}

func TestSample3(t *testing.T) {
	runSample(t, "MMFF", 3)
}

func TestSample4(t *testing.T) {
	runSample(t, "MFFFMMFMFMFMFFFMMMFFMMMMMMFMMFFMMMFMMFMFFFMMFMMMFFMMFFFFFMFMFFFMMMFFFMFMFMFMFFFMMMMFMMFMMFFMMMMMMFFM", 54)
}
