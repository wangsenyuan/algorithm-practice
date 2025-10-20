package p3720

import "testing"

func runSample(t *testing.T, s string, target string, expect string) {
	res := lexGreaterPermutation(s, target)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abc"
	target := "bba"
	expect := "bca"
	runSample(t, s, target, expect)
}

func TestSample2(t *testing.T) {
	s := "leet"
	target := "code"
	expect := "eelt"
	runSample(t, s, target, expect)
}

func TestSample3(t *testing.T) {
	s := "baba"
	target := "bbaa"
	expect := ""
	runSample(t, s, target, expect)
}

func TestSample4(t *testing.T) {
	s := "aab"
	target := "abb"
	expect := "baa"
	runSample(t, s, target, expect)
}
