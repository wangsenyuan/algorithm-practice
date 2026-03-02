package main

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{1, 1, 0},
		{2, 2, 8},
	}
	for _, tt := range tests {
		got := solve(tt.a, tt.b)
		if got != tt.expect {
			t.Errorf("solve(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expect)
		}
	}
}
