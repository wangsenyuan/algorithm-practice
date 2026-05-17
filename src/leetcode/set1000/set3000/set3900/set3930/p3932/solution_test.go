package p3932

import "testing"

func runSample(t *testing.T, l int, r int, k int, expect int) {
	res := countKthRoots(l, r, k)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 16, 16, 2, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 9, 3, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 45, 9, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, 0, 0, 30, 1)
}

