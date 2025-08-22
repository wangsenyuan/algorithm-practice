package main

import "testing"

func runSample(t *testing.T, x int, y int, z int, k int, expect int) {
	res := solve(x, y, z, k)

	if res != expect {
		t.Errorf("Sample %d, %d, %d, %d, expect %d, but got %d", x, y, z, k, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 2, 2, 3, 8)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2, 2, 1, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 1, 1, 1, 1, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 100, 500, 100500, 1000000000, 5025000000)
}

func TestSample5(t *testing.T) {
	runSample(t, 999999, 1, 999998, 1333333, 444445555556)
}
