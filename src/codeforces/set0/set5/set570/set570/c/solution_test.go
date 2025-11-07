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
	s := `10 3
.b..bz....
1 h
3 c
9 f`
	expect := []int{4, 3, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
.cc.
2 .
3 .
2 a
1 a`
	expect := []int{1, 3, 1, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1 5
.
1 .
1 w
1 w
1 .
1 .`
	expect := []int{0, 0, 0, 0, 0}
	runSample(t, s, expect)
}
