package main

import "testing"

func runSample(t *testing.T, m int, expect int) {
	ans := solve(m)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 7, 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 8, 1440)
}

func TestSample3(t *testing.T) {
	runSample(t, 47, 907362803)
}

func TestSample4(t *testing.T) {
	runSample(t, 128, 879893164)
}
