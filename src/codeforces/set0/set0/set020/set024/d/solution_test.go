package main

import "testing"

func TestSample1(t *testing.T) {
	got := solve(10, 10, 10, 4)
	if got != 0 {
		t.Fatalf("expect 0, got %.10f", got)
	}
}

func TestSample2(t *testing.T) {
	got := solve(10, 14, 5, 14)
	want := 18.0038068653
	if diff(got, want) > 1e-8 {
		t.Fatalf("expect %.10f, got %.10f", want, got)
	}
}

func TestSingleColumn(t *testing.T) {
	got := solve(7, 1, 3, 1)
	want := 8.0
	if diff(got, want) > 1e-9 {
		t.Fatalf("expect %.10f, got %.10f", want, got)
	}
}

func diff(a, b float64) float64 {
	if a > b {
		return a - b
	}
	return b - a
}
