package main

import "testing"

func runSample(t *testing.T, r, g int, expect int) {
	ans := solve(r, g)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 6, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 7, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1, 2)
}

// func TestSample4(t *testing.T) {
// 	runSample(t, 200000, 200000, 206874596)
// }
