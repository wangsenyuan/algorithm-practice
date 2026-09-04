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
	runSample(t, `16
-5 1 1 2 2 3 3 4 4 5 5 6 6 7 7 10
`, []int{1, 2, 1, 7, 6, 2, 6, 7})
}

func TestSample2(t *testing.T) {
	runSample(t, `8
0 0 -1 2 2 1 1 3
`, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, `8
0 0 0 0 0 5 0 5
`, []int{0, 0, 0, 5, 0, 0, 0, 5})
}
