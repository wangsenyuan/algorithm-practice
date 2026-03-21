package main

import "testing"

func runSample(t *testing.T, r int, g int, b int, expect int) {
	res := solve(r, g, b)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 6, 9, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 4, 4, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 0, 0, 0, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 7, 8, 9, 7)
}

func TestSample8(t *testing.T) {
	runSample(t, 8, 8, 9, 8)
}

func TestSample9(t *testing.T) {
	runSample(t, 3, 5, 5, 4)
}
