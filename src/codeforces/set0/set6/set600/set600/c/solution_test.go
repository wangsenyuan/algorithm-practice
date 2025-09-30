package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aabc"
	expect := "abba"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aabcd"
	expect := "abcba"
	runSample(t, s, expect)
}	

func TestSample3(t *testing.T) {
	s := "aabbcccdd"
	expect := "abcdcdcba"
	runSample(t, s, expect)
}	