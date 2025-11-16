package p3747

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := countDistinct(n)
	if res != expect {
		t.Errorf("Sample failed, expect %d, got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 10, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 3)
}