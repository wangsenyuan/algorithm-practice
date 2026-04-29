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
	s := `11
5 5 5 5 2 2 2 8 6 1 7
`
// 1 2 2 2 5 5 5 5 6 7 8
	expect := []int{3, 3, 2, 2, 2, 1, 1, 1, 1, 1, 1}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
3 3 3 3 3 3
`
	expect := []int{1, 1, 1, 1, 1, 1}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `5
2 1 3 5 4
`
	expect := []int{1, 1, 1, 1, 1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8
1 1 1 2 3 4 5 6
`
	expect := []int{2, 2, 1, 1, 1, 1, 1, 1}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `1
1
`
	expect := []int{1}
	runSample(t, s, expect)
}
