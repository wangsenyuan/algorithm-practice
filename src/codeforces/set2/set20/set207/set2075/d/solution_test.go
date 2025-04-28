package main

import "testing"

func runSample(t *testing.T, x int, y int, expect int) {
	res := solve(x, y)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, 1, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 2, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 13, 37, 26)
}

func TestSample4(t *testing.T) {
	runSample(t, 4238659325782394, 12983091057341925, 32764)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 3, 12)
}
