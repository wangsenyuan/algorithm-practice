package main

import "testing"

func runSample(t *testing.T, x string, expect int) {
	res := solve(x)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 0, 1, 2, 3
	// 3, 2, 1, 0
	// 00   11
	// 01   10
	// 10   01
	// 11   00
	runSample(t, "11", 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "01", 2)
}

func TestSample3(t *testing.T) {
	// 0 1
	// 1,0
	runSample(t, "1", 1)
}
