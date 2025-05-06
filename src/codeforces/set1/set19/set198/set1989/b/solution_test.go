package main

import "testing"

func runSample(t *testing.T, a, b string, expect int) {
	res := solve(a, b)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := "aba"
	b := "cb"
	expect := 4
	runSample(t, a, b, expect)
}

func TestSample2(t *testing.T) {
	a := "er"
	b := "cf"
	expect := 4
	runSample(t, a, b, expect)
}

func TestSample3(t *testing.T) {
	a := "mmm"
	b := "mmm"
	expect := 3
	runSample(t, a, b, expect)
}

func TestSample4(t *testing.T) {
	a := "contest"
	b := "test"
	expect := 7
	runSample(t, a, b, expect)
}

func TestSample5(t *testing.T) {
	a := "cde"
	b := "abcefg"
	expect := 7
	runSample(t, a, b, expect)
}