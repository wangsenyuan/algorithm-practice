package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 4, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 7, 1)
}

func TestSample3(t *testing.T) {
	// 20
	runSample(t, 20, 1, 2)
}

func TestSample4(t *testing.T) {
	// 4, 7, 44, 77, 47, 74
	runSample(t, 100, 1, 6)
}

func TestSample5(t *testing.T) {
	// 4, 7, 44, 77, 47, 74, 444, 447, 474, 477, 744, 747, 774, 777
	runSample(t, 778, 1, bruteForce(778, 1))
}
