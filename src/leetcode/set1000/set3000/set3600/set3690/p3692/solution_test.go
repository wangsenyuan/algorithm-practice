package p3692

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := majorityFrequencyGroup(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaabbbccdddde"
	expect := "ab"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abcd"
	expect := "abcd"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "pfpfgi"
	expect := "fp"
	runSample(t, s, expect)
}
