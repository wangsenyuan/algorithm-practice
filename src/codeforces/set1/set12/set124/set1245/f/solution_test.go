package main

import "testing"

func runSample(t *testing.T, L int, R int, expect int) {
	res := solve(L, R)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 4, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 323, 323, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1000000, 3439863766)
}