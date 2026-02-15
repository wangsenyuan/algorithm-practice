package p3844

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := almostPalindromic(s)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "abba"
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "abca"
	expect := 4
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "zzabba"
	expect := 5
	runSample(t, s, expect)
}
