package p3966

import "testing"

func runSample(t *testing.T, l int64, r int64, k int, expect int64) {
	t.Helper()
	res := goodIntegers(l, r, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 15, 1, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 201, 204, 2, 2)
}
