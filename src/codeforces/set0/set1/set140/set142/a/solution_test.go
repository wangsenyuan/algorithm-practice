package main

import "testing"

func runSample(t *testing.T, n int, expect []int) {
	res := solve(n)

	if res[0] != expect[0] || res[1] != expect[1] {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, []int{28, 41})
}

func TestSample2(t *testing.T) {
	runSample(t, 7, []int{47, 65})
}

func TestSample3(t *testing.T) {
	runSample(t, 12, []int{48, 105})
}

func TestSample4(t *testing.T) {
	runSample(t, 299999771, []int{1499998867, 2399998177})
}
