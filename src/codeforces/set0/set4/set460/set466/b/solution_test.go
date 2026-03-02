package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectArea, expectA, expectB int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	area, a1, b1 := drive(reader)
	if area != expectArea {
		t.Errorf("area = %d, want %d", area, expectArea)
	}
	if a1 != expectA || b1 != expectB {
		t.Errorf("sides = %d %d, want %d %d", a1, b1, expectA, expectB)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "3 3 5\n", 18, 3, 6)
}

func TestSample2(t *testing.T) {
	runSample(t, "2 4 4\n", 16, 4, 4)
}

func TestSample3(t *testing.T) {
	runSample(t, "8 7 5\n", 48, 8, 6)
}

func TestSolve(t *testing.T) {
	tests := []struct {
		n, a, b    int
		expectArea int
	}{
		// area already sufficient
		{1, 3, 3, 9},
		// extend b only
		{5, 2, 3, 30},
		// extend a only
		{5, 3, 2, 30},
	}
	for _, tt := range tests {
		res := solve(tt.n, tt.a, tt.b)
		area, a1, b1 := res[0], res[1], res[2]
		if area != tt.expectArea {
			t.Errorf("solve(%d,%d,%d): area=%d, want %d", tt.n, tt.a, tt.b, area, tt.expectArea)
		}
		if a1 < tt.a || b1 < tt.b {
			t.Errorf("solve(%d,%d,%d): sides (%d,%d) smaller than original (%d,%d)", tt.n, tt.a, tt.b, a1, b1, tt.a, tt.b)
		}
		if a1*b1 < 6*tt.n {
			t.Errorf("solve(%d,%d,%d): area %d < required %d", tt.n, tt.a, tt.b, a1*b1, 6*tt.n)
		}
		if a1*b1 != area {
			t.Errorf("solve(%d,%d,%d): reported area %d != actual %d", tt.n, tt.a, tt.b, area, a1*b1)
		}
	}
}
