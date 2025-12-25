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
	s := `4
1 3 2 5
3 4 1 2
`
	expect := []int{5, 4, 3, 0}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2 3 4 5
4 2 3 5 1
`
	expect := []int{6, 5, 5, 1, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8
5 5 4 4 6 6 5 5
5 2 8 7 1 3 4 6
`
	expect := []int{18, 16, 11, 8, 8, 6, 6, 0}
	runSample(t, s, expect)
}
