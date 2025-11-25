package main

import "testing"

func runSample(t *testing.T, n int, expect bool) {
	res := solve(n)
	if res != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 8, false)
}

func TestSample5(t *testing.T) {
	runSample(t, 11, false)
}