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
	s := `2 1 3
0 1
2
LRR`
	expect := []int{2, 2, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3 3
2 4
1 3 5
LRL`
	expect := []int{0, 0, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 2 3
1 3 7
9 6
RRL`
	expect := []int{3, 2, 2}
	runSample(t, s, expect)
}
