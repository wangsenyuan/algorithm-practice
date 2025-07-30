package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 2, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 1, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 570441179141911871, 511467058318039545, 511467058661475480)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 1, 2)
}
