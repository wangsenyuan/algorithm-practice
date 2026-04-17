package main

import "testing"

func TestSample1(t *testing.T) {
	cases := []struct {
		a, b int
		ok   bool
	}{
		{2, 4, true},
		{75, 45, true},
		{8, 8, true},
		{16, 16, false},
		{247, 994, false},
		{1000000000, 1000000, true},
	}

	for _, tc := range cases {
		if got := solve(tc.a, tc.b); got != tc.ok {
			t.Fatalf("solve(%d, %d) expect %v, got %v", tc.a, tc.b, tc.ok, got)
		}
	}
}
