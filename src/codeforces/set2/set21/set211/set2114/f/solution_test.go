package main

import "testing"

func runSample(t *testing.T, x, y, k int, expect int) {
	res := solve(x, y, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 6, 3, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 5, 3, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 6, 2, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 780, 23, 42, 3)
}

func TestSample5(t *testing.T) {
	runSample(t, 11, 270, 23, 3)
}

func TestSample6(t *testing.T) {
	runSample(t, 1, 982800, 13, 6)
}

func TestSample7(t *testing.T) {
	runSample(t, 1, 6, 2, -1)
}
