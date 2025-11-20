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
	s := `3 4 2 2
UURDRDRL
`
	expect := []int{1, 1, 0, 1, 1, 1, 1, 0, 6}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 2 2 2
ULD
`
	expect := []int{1, 1, 1, 1}
	runSample(t, s, expect)
}
