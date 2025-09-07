package p3677

import "testing"

func runSample(t *testing.T, n int64, expect int) {
	res := countBinaryPalindromes(n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 9, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, 0, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1000000000000000, 63356754)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 2)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 2)
}