package p3791

import "testing"

func runSample(t *testing.T, low int, high int, expect int) {
	res := countBalanced(int64(low), int64(high))
	if res != int64(expect) {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 100, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 120, 129, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 1234, 1234, 0)
}