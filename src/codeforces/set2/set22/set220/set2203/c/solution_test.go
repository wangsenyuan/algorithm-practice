package main

import "testing"

func runSample(t *testing.T, s int, m int, expect int) {
	res := solve(s, m)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 13, 5, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 3, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, 13, 6, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000000007, 2776648, -1)
}

func TestSample5(t *testing.T) {
	runSample(t, 99999999999, 1, 99999999999)
}

func TestSample6(t *testing.T) {
	runSample(t, 998244353, 1557287, 642)
}
