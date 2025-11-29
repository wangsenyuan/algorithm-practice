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
	s := `3 1
1 2 3
3
2 1 3
1 1 2 4
2 1 3
`
	expect := []int{1, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 2
1 2 3
5
2 4 4
1 4 4 5
2 4 4
1 1 6 1
2 6 6
`
	expect := []int{1, 5, 1}
	runSample(t, s, expect)
}
