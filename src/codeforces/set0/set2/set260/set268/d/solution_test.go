package main

import "testing"

func runSample(t *testing.T, n int, h int, expect int) {
	res := solve(n, h)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 1, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 2, 148)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 3, 256)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 2, 376)
}

// func TestSample5(t *testing.T) {
// 	runSample(t, 1000, 30, 107282225)
// }
