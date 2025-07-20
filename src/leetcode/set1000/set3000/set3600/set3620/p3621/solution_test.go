package p3621

import "testing"

func runSample(t *testing.T, n int64, k int, expect int64) {
	res := popcountDepth(n, k)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, 1, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, 7, 2, 3)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 1, 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 1000000000, 3, 486137281)
}
