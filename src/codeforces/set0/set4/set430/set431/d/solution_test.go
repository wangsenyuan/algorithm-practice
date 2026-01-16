package main

import "testing"

func runSample(t *testing.T, m int, k int) {
	n := solve(m, k)

	w := count(n*2, k) - count(n, k)

	if w != m {
		t.Fatalf("Sample expect %d, but got %d", m, w)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 1, 1)
}

func TestSample2(t *testing.T) {
	// 1, 2, 3, 4, 5, 6
	//       2     2  2
	// 8， 9 10 12， 13， 14
	runSample(t, 3, 2)
}

func TestSample3(t *testing.T) {
	// 1, 2, 3, 4, 5, 6
	//       2     2  2
	// 8， 9 10 12， 13， 14
	runSample(t, 3, 3)
}
