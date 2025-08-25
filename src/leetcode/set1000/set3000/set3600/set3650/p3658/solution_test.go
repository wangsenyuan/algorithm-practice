package p3658

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := gcdOfOddEvenSums(n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 4)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 5)
}