package p3770

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := largestPrime(n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 20, 17)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 2)
}
