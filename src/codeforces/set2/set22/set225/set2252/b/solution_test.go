package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "0101"
	expect := 0
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "111"
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "100110"
	expect := 2
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "100010"
	expect := 3
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "011110"
	expect := 5
	runSample(t, s, expect)
}
