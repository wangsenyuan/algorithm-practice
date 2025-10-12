package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4
DRLD
U.UL
.UUR
RDDL
`
	expect := []int{10, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 5
.D...
RRRLL
.U...
`
	expect := []int{6, 2}
	runSample(t, s, expect)
}
