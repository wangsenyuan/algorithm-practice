package p3628

import "testing"

func runSample(t *testing.T, s string, expect int64) {
	res := numOfSubsequences(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "LMCT", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "LCCT", 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "L", 0)
}

func TestSample4(t *testing.T) {
	runSample(t, "LCLPTTGC", 6)
}
