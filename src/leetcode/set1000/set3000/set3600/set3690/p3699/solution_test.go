package p3699

import "testing"

func runSample(t *testing.T, n int, l int, r int, expect int) {
	res := zigZagArrays(n, l, r)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 4, 5, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 1, 3, 10)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 2, 4, 16)
}

func TestSample4(t *testing.T) {
	runSample(t, 2920, 2, 18, 604859794)
}
