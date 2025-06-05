package main

import "testing"

func runSample(t *testing.T, n int, m int, s int, expect int) {
	res := solve(n, m, s)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 1, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 4, 5, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 3, 5, 2)
}

func TestSample4(t *testing.T) {
	// 00000
	// ###00, 0###0, 00###
	// (0,1,0,1,0,1)
	//
	runSample(t, 5, 1, 3, 9)
}

func TestSample5(t *testing.T) {
	runSample(t, 5, 10, 25, 102)
}

func TestSample6(t *testing.T) {
	runSample(t, 20, 12, 101, 424)
}
