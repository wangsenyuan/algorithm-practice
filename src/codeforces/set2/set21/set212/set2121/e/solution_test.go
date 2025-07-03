package main

import "testing"

func runSample(t *testing.T, l string, r string, expect int) {
	res := solve(l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "123", "456", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, "1", "1", 2)
}

func TestSample3(t *testing.T) {
	runSample(t, "2", "3", 1)
}

func TestSample4(t *testing.T) {
	runSample(t, "4", "6", 0)
}

func TestSample5(t *testing.T) {
	runSample(t, "15", "16", 3)
}

func TestSample6(t *testing.T) {
	runSample(t, "17", "19", 2)
}

func TestSample7(t *testing.T) {
	runSample(t, "199", "201", 2)
}

func TestSample8(t *testing.T) {
	runSample(t, "899", "999", 1)
}

func TestSample9(t *testing.T) {
	runSample(t, "1990", "2001", 3)
}

func TestSample10(t *testing.T) {
	runSample(t, "6309", "6409", 3)
}

func TestSample11(t *testing.T) {
	runSample(t, "12345", "12501", 4)
}

func TestSample12(t *testing.T) {
	runSample(t, "19987", "20093", 3)
}

func TestSample13(t *testing.T) {
	runSample(t, "746814", "747932", 5)
}

func TestSample14(t *testing.T) {
	runSample(t, "900990999", "900991010", 12)
}

func TestSample15(t *testing.T) {
	runSample(t, "999999999", "999999999", 18)
}
