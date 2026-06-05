package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ACBBCABCAB"
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "CABACBBBBBAABABACCBCABCCABAABABBCBAC"
	expect := 136
	runSample(t, s, expect)
}
