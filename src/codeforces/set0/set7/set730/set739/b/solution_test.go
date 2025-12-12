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
	s := `5
2 5 1 4 6
1 7
1 1
3 5
3 6
	`
	expect := []int{1, 0, 1, 0, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
9 7 8 6 5
1 1
2 1
3 1
4 1
	`
	expect := []int{4, 3, 2, 1, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
96 92 63 25 80 74 95 41 28 54
6 98
1 11
5 45
3 12
7 63
4 39
7 31
8 81
2 59
	`
	expect := []int{2, 0, 1, 1, 1, 0, 2, 0, 0, 0}
	runSample(t, s, expect)
}
