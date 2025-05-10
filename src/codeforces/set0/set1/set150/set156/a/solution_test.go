package main

import "testing"

func runSample(t *testing.T, s string, u string, expect int) {
	res := solve(s, u)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "aaaaa"
	u := "aaa"
	expect := 0
	runSample(t, s, u, expect)
}

func TestSample2(t *testing.T) {
	s := "abcabc"
	u := "bcd"
	expect := 1
	runSample(t, s, u, expect)
}

func TestSample3(t *testing.T) {
	s := "abcdef"
	u := "klmnopq"
	expect := 7
	runSample(t, s, u, expect)
}

func TestSample4(t *testing.T) {
	s := "aaabbbaaa"
	u := "aba"
	expect := 1
	runSample(t, s, u, expect)
}

func TestSample5(t *testing.T) {
	s := "aaabbbaaa"
	u := "abcba"
	expect := 1
	runSample(t, s, u, expect)
}