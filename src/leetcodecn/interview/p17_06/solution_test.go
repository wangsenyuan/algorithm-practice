package p1706

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := numberOf2sInRange(n)
	if res != expect {
		t.Errorf("Sample %d, expect %d, but got %d", n, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 25
	expect := 9
	runSample(t, n, expect)
}

func TestSample2(t *testing.T) {
	n := 2
	expect := 1
	runSample(t, n, expect)
}

func TestSample3(t *testing.T) {
	n := 100000000
	expect := 80000000

	runSample(t, n, expect)
}
