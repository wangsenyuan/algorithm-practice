package main

import "testing"

func runSample(t *testing.T, num int, expect int) {
	res := solve(num)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 16, 4)
}
