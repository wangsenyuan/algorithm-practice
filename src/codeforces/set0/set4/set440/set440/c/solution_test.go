package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}
func TestSample1(t *testing.T) {
	runSample(t, 121, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 81924761239462, 321)
}

func TestSample3(t *testing.T) {
	runSample(t, 2038946593, 145)
}
