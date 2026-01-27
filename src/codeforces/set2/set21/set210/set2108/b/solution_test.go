package main

import "testing"

func runSample(t *testing.T, n int, x int, expect int) {
	res := solve(n, x)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 6, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 0, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 0, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, 5, 0, 8)
}

func TestSample6(t *testing.T) {
	runSample(t, 2, 27, 27)
}

func TestSample7(t *testing.T) {
	runSample(t, 15, 43, 55)
}

func TestSample8(t *testing.T) {
	runSample(t, 12345678, 9101112, 21446778)
}
