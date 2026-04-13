package main

import "testing"

func runSample(t *testing.T, x int, y int, expect int) {
	res := solve(x, y)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	x, y := 6, 3
	expect := 4
	runSample(t, x, y, expect)
}

func TestSample2(t *testing.T) {
	x, y := 8, 5
	expect := 3
	runSample(t, x, y, expect)
}

func TestSample3(t *testing.T) {
	x, y := 22, 4
	expect := 6
	runSample(t, x, y, expect)
}
