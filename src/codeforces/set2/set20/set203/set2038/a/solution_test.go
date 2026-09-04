package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 6
4 7 6
1 2 3
`, []int{1, 3, 2})
}

func TestSample2(t *testing.T) {
	runSample(t, `3 12
4 7 6
1 2 3
`, []int{0, 0, 0})
}

func TestSample3(t *testing.T) {
	runSample(t, `3 11
6 7 8
1 2 3
`, []int{6, 3, 2})
}
