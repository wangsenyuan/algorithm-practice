package main

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := solve(n, x)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 2, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000000, 1048576, 118606527258)
}

func TestSample4(t *testing.T) {
	runSample(t, 82426873, 1, 26)
}

func TestSample5(t *testing.T) {
	runSample(t, 4890852, 16, 31009)
}
