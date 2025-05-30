package main

import "testing"

func runSample(t *testing.T, n int, m int, k int, p int, expect int) {
	res := solve(n, m, k, p)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 3, 2, 10007, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 4, 1, 769626776, 472)
}

func TestSample3(t *testing.T) {
	runSample(t, 66, 11, 9, 786747482, 206331312)
}

func TestSample4(t *testing.T) {
	runSample(t, 99, 30, 18, 650457567, 77365367)
}
