package main

import "testing"

func runSample(t *testing.T, n int, m int, expect int) {
	res := solve(n, m)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 2, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 123, 456, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 20260530, 460, 922576091)
}

func TestSample4(t *testing.T) {
	runSample(t, 123456789123456789, 998244353, 422081792)
}

func TestSample5(t *testing.T) {
	runSample(t, 1000000000000000000, 2, 715642751)
}
