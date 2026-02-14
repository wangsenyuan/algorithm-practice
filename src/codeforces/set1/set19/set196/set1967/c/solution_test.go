package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)

	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `8 1
1 2 1 4 1 2 1 8`
	expect := []int{1, 1, 1, 1, 1, 1, 1, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6 2
1 4 3 17 5 16`
	expect := []int{1, 2, 3, 4, 5, 6}
	runSample(t, s, expect)
}
