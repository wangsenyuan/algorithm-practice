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
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2
11 8
`
	expect := []int{9, 10}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
10 9 7 10 6
`
	// expect := []int{9, 10}
	runSample(t, s, nil)
}

func TestSample3(t *testing.T) {
	s := `3
12 3 3
`
	expect := []int{4, 4, 10}
	runSample(t, s, expect)
}
