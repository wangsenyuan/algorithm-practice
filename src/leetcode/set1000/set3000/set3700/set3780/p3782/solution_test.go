package p3782

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := lastInteger(n)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 8, 3)
}

func TestSample2(t *testing.T) {
	runSample(t, 5, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 9, 9)
}
