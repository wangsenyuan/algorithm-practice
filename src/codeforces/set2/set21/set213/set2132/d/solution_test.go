package main

import (
	"testing"
)

func runSample(t *testing.T, k int, expect int) {
	res := solve(k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	// 12345
	runSample(t, 5, 15)
}

func TestSample2(t *testing.T) {
	runSample(t, 10, 46)
}

func TestSample3(t *testing.T) {
	runSample(t, 13, 48)
}

func TestSample4(t *testing.T) {
	runSample(t, 29, 100)
}

func TestSample5(t *testing.T) {
	runSample(t, 1000000000, 4366712386)
}

func TestSample6(t *testing.T) {
	runSample(t, 1000000000000000, 4441049382716054)
}
