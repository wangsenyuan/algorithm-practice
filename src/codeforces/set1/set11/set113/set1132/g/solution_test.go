package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 4
1 5 2 5 3 6
`, []int{2, 2, 3})
}

func TestSample2(t *testing.T) {
	runSample(t, `7 6
4 5 2 5 3 6 6
`, []int{3, 3})
}
