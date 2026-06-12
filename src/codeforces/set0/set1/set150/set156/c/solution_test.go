package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "ab"
	expect := 1
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aaaaaaaaaaa"
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "ya"
	expect := 24
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "klmbfxzb"
	expect := 320092793
	runSample(t, s, expect)
}
