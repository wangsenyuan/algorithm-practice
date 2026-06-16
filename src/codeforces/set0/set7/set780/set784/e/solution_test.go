package main

import "testing"

func runSample(t *testing.T, a int, b int, c int, d int, expect int) {
	res := solve(a, b, c, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 0, 1, 1, 0, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 0, 0, 0, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 0, 0, 0, 0, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 0, 0, 1, 1)
}
