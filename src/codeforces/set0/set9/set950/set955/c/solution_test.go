package main

import "testing"

func runSample(t *testing.T, l int, r int, expect int) {
	res := solve(l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 4, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 9, 9, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 5, 7, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 12, 29, 3)
}

func TestSample5(t *testing.T) {
	runSample(t, 137, 591, 17)
}

func TestSample6(t *testing.T) {
	runSample(t, 1, 1000000, 1111)
}

func TestSample7(t *testing.T) {
	runSample(t, 387291074607832779, 798305191127761550, 271351299)
}


func TestSample8(t *testing.T) {
	runSample(t, 2, 999999999999999999, 1001003330)
}
