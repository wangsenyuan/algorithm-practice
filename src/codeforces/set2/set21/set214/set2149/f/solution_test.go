package main

import "testing"

func runSample(t *testing.T, hp int, d int, expect int) {
	res := solve(hp, d)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 1, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 3, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, 2, 4, 7)
}

func TestSample5(t *testing.T) {
	runSample(t, 10, 7, 10)
}
