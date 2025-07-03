package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}
func TestSample2(t *testing.T) {
	runSample(t, 4, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 8, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 10, 4)
}

func TestSample5(t *testing.T) {
	runSample(t, 99999999999999999, 3703703703703704)
}
