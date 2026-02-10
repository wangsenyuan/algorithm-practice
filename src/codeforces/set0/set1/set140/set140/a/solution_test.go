package main

import "testing"

func runSample(t *testing.T, n int, R int, r int, expect bool) {
	res := solve(n, R, r)
	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	R := 10
	r := 4
	expect := true
	runSample(t, n, R, r, expect)
}

func TestSample2(t *testing.T) {
	n := 5
	R := 10
	r := 4
	expect := false
	runSample(t, n, R, r, expect)
}

func TestSample3(t *testing.T) {
	n := 1
	R := 10
	r := 10
	expect := true
	runSample(t, n, R, r, expect)
}