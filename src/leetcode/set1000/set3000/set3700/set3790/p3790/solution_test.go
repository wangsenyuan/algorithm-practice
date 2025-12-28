package p3790

import "testing"

func runSample(t *testing.T, k int, expect int) {
	res := minAllOneMultiple(k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, -1)
}