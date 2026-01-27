package main

import "testing"

func runSample(t *testing.T, x int, y int, l int, r int, expect int) {
	res := solve(x, y, l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 3, 1, 10, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 5, 10, 22, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 3, 3, 5, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 2, 1, 1000000, 213568)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 14, 732028847861235712, 732028847861235712, 0)
}

func TestSample6(t *testing.T) {
	runSample(t, 3, 3, 1, 1, 1)
}
