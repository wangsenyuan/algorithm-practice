package p3557

import "testing"

func runSample(t *testing.T, word string, expect int) {
	res := maxSubstrings(word)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "abcdeafdef", 2)
}

func TestSample2(t *testing.T) {
	runSample(t, "bcdaaaab", 1)
}
