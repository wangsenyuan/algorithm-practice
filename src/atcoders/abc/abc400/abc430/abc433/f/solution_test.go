package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "1122"
	expect := 5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "0777468889971"
	expect := 30
	runSample(t, s, expect)
}
