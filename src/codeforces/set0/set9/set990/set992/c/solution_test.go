package main

import "testing"

func runSample(t *testing.T, x int, k int, expect int) {
	res := solve(x, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 0, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, 7)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 2, 21)
}

func TestSample4(t *testing.T) {
	runSample(t, 348612312017571993, 87570063840727716, 551271547)
}
