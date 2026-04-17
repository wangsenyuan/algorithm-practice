package main

import "testing"

func runSample(t *testing.T, mask string, expect string) {
	t.Helper()
	best := precompute()
	m := 0
	for i := 0; i < 8; i++ {
		if mask[i] == '1' {
			m |= 1 << i
		}
	}
	if best[m] != expect {
		t.Fatalf("expect %s, got %s", expect, best[m])
	}
}

func TestSampleMasks(t *testing.T) {
	runSample(t, "00110011", "y")
	runSample(t, "11110000", "!x")
	runSample(t, "00011111", "x|y&z")
}

func TestSample2(t *testing.T) {
	runSample(t, "00000111", "(y|z)&x")
}
