package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int64) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
202516
`, []int64{4, 8, 10, 10, 9, 6})
}

func TestSample2(t *testing.T) {
	runSample(t, `5
20265
`, []int64{5, 8, 9, 8, 5})
}
