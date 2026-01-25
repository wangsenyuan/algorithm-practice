package p3821

import "testing"

func runSample(t *testing.T, n int64, k int, expect int64) {
	res := nthSmallest(n, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 2, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 1, 4)
}
