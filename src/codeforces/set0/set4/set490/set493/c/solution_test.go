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
	s := `3
1 2 3
2
5 6
`
	expect := []int{9, 6}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
6 7 8 9 10
5
1 2 3 4 5
`
	expect := []int{15, 10}
	runSample(t, s, expect)
}
