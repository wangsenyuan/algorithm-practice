package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 6)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, 10, 12)
}

func TestSample6(t *testing.T) {
	runSample(t, 9, 10)
}
