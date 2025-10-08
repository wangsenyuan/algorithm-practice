package p3703

import "testing"

func runSample(t *testing.T, n int64, expect int64) {
	res := countNoZeroPairs(n)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 2, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 11, 8)
}

func TestSample3(t *testing.T) {
	runSample(t, 3, 2)
}

func TestSample4(t *testing.T) {
	// 9 + 12
	// 8 + 13
	// 7 + 14
	// 6 + 15
	// 5 + 16
	// 4 + 17
	// 3 + 18
	// 2 + 19
	runSample(t, 1000, bruteForce(1000))
}
