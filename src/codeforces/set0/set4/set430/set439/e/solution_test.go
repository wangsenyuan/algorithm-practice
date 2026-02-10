package main

import (
	"testing"
)

func runSample(t *testing.T, n, f int, expect int) {
	res := solve(n, f)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 6, 2, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 2, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 6, 3, 9)
}

func TestSample4(t *testing.T) {
	runSample(t, 6, 4, 10)
}

func TestSample5(t *testing.T) {
	runSample(t, 7, 4, 20)
}

func TestEdge1(t *testing.T) {
	runSample(t, 1, 1, 1)
}

func TestEdge2(t *testing.T) {
	runSample(t, 10, 1, 0)
}

func TestEdge3(t *testing.T) {
	runSample(t, 4, 1, 0)
}

func TestEdge4(t *testing.T) {
	runSample(t, 32, 2, 16)
}
