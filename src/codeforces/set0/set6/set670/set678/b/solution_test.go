package main

import "testing"

func runSample(t *testing.T, y int, expect int) {
	res := solve(y)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2016, 2044)
}

func TestSample2(t *testing.T) {
	runSample(t, 2000, 2028)
}

func TestSample3(t *testing.T) {
	runSample(t, 50501, 50507)
}
