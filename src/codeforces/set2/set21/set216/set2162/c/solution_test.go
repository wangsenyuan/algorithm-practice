package main

import "testing"

func runSample(t *testing.T, a int, b int, expect bool) {
	ok, res := solve(a, b)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}
	w := a
	for _, v := range res {
		if v > a {
			t.Fatalf("Sample result %v, but got %d > %d", res, v, a)
		}
		w ^= v
	}

	if w != b {
		t.Fatalf("Sample result %v, but got %d != %d", res, w, b)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 9, 6, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 13, 13, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 292, 929, false)
}

func TestSample4(t *testing.T) {
	runSample(t, 405, 400, true)
}

func TestSample5(t *testing.T) {
	runSample(t, 998, 244, true)
}

func TestSample6(t *testing.T) {
	runSample(t, 244, 353, false)
}

func TestSample7(t *testing.T) {
	runSample(t, 28, 31, true)
}
