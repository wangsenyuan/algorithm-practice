package main

import "testing"

func runSample(t *testing.T, a int, b int, expect int) {
	res := solve(a, b)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 5, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 6, 3, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000009, 1000000008, 1000000008)
}

func TestSample4(t *testing.T) {
	runSample(t, 2000000018, 2000000017, 1000000009)
}

func TestSample6(t *testing.T) {
	runSample(t, 1000000000000, 1, 1)
}

func TestSample7(t *testing.T) {
	runSample(t, 1, 1000000000000, 1000000000000)
}

func TestSample5(t *testing.T) {
	runSample(t, 49544527863, 318162327511, 6965053451)
}
