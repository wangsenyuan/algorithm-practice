package main

import "testing"

func validate(t *testing.T, res []int, a, b, m int) {
	t.Helper()
	if res[0] != a || res[len(res)-1] != b {
		t.Fatalf("Expected first=%d last=%d, got first=%d last=%d", a, b, res[0], res[len(res)-1])
	}
	s := res[0]
	for i := 1; i < len(res); i++ {
		r := res[i] - s
		if r < 1 || r > m {
			t.Fatalf("At position %d: r=%d not in [1,%d], seq=%v", i, r, m, res)
		}
		s += res[i]
	}
}

func runSample(t *testing.T, a, b, m int, possible bool) {
	res := solve(a, b, m)
	if !possible {
		if res != nil {
			t.Fatalf("Expected impossible, but got %v", res)
		}
		return
	}
	if res == nil {
		t.Fatalf("Expected possible for a=%d b=%d m=%d, but got impossible", a, b, m)
	}
	validate(t, res, a, b, m)
}

func TestSample1(t *testing.T) {
	runSample(t, 5, 26, 2, true)
}

func TestSample2(t *testing.T) {
	runSample(t, 3, 9, 1, false)
}

func TestSingleElement(t *testing.T) {
	runSample(t, 7, 7, 3, true)
}

func TestTwoElements(t *testing.T) {
	runSample(t, 1, 2, 1, true)
}

func TestSmallImpossible(t *testing.T) {
	// m=1: sequence 1,2,4,8,... so b=5 is unreachable
	runSample(t, 1, 5, 1, false)
}

func TestLargeM(t *testing.T) {
	// m >= b, so [a, b] always works
	runSample(t, 3, 10, 100, true)
}

func TestSample4(t *testing.T) {
	// m >= b, so [a, b] always works
	runSample(t, 1, 1350, 3, true)
}
