package main

import (
	"testing"
)

func runSample(t *testing.T, n int64, expectOK bool) {
	t.Helper()
	x, y, ok := solve(n)
	if ok != expectOK {
		t.Errorf("Sample expect ok=%v, but got ok=%v (%d %d)", expectOK, ok, x, y)
		return
	}
	if !expectOK {
		return
	}
	w := cubic(int64(x)) - cubic(int64(y))
	if w != n {
		t.Errorf("Sample result %d - %d, got %d, but want %d", x, y, w, n)
		return
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 397, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, false)
}

func TestSample3(t *testing.T) {
	runSample(t, 39977273855577088, true)
}
