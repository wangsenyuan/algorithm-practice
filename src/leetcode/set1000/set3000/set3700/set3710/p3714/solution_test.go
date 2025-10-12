package p3714

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := longestBalanced(s)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abbac", 4)
}

func TestSample2(t *testing.T) {
	runSample(t, "aabcc", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "aba", 2)
}
