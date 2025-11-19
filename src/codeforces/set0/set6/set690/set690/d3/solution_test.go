package main

import "testing"

func runSample(t *testing.T, C int, W int, H int, expect int) {
	res := solve(C, W, H)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	C := 1
	W := 1
	H := 1
	expect := 2
	runSample(t, C, W, H, expect)
}

func TestSample2(t *testing.T) {
	C := 1
	W := 2
	H := 2
	expect := 3
	runSample(t, C, W, H, expect)
}

func TestSample3(t *testing.T) {
	C := 1
	W := 2
	H := 3
	expect := 4
	runSample(t, C, W, H, expect)
}

func TestSample4(t *testing.T) {
	C := 3
	W := 2
	H := 2
	expect := 19
	runSample(t, C, W, H, expect)
}

func TestSample5(t *testing.T) {
	C := 5
	W := 4
	H := 9
	expect := 40951
	runSample(t, C, W, H, expect)
}

func TestSample6(t *testing.T) {
	C := 40
	W := 37
	H := 65
	expect := 933869
	runSample(t, C, W, H, expect)
}
