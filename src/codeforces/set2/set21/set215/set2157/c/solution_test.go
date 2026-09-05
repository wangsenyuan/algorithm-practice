package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	t.Skip("solve TODO")
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 2 2
1 1 3
2 2 6
`, []int{2, 5, 4, 3, 0, 1})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3 1
2 1 3
`, []int{2, 0, 1})
}

func TestSample3(t *testing.T) {
	runSample(t, `3 3 2
1 1 1
1 3 3
`, []int{3, 3, 3})
}

func TestSample4(t *testing.T) {
	runSample(t, `3 2 2
2 1 2
2 2 3
`, []int{1, 0, 1})
}
