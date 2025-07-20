package main

import "testing"

func runSample(t *testing.T, s string, expect int) {
	res := solve(s)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "0", 1)
}
func TestSample2(t *testing.T) {
	runSample(t, "01", 3)
}

func TestSample3(t *testing.T) {
	runSample(t, "0110", 14)
}

func TestSample4(t *testing.T) {
	runSample(t, "110001", 40)
}

func TestSample5(t *testing.T) {
	runSample(t, "10011100", 78)
}

func TestSample6(t *testing.T) {
	runSample(t, "01011011100", 190)
}
