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
	s := `2 8
12 20
`
	expect := []int{0, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 10
10 20 30
`
	expect := []int{0}
	runSample(t, s, expect)
}
