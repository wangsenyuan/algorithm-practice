package main

import "testing"

func runSample(t *testing.T, n int, k int, expect int) {
	res := solve(n, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 1, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 2, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, 42, 2, 3)
}

func TestSample4(t *testing.T) {
	runSample(t, 1048576, 100, 100)
}

func TestSample5(t *testing.T) {
	runSample(t, 23, 2, 5)
}

func TestSample6(t *testing.T) {
	runSample(t, 371, 1, 3)
}


func TestSample7(t *testing.T) {
	runSample(t, 413805207, 9, 17)
}

func TestSample8(t *testing.T) {
	runSample(t, 1, 2, 2)
}
