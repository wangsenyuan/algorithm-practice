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
	s := `8
2 4 3 6 5 7 8 6`
	expect := []int{0, 1, 2, 3, 3, 3, 4, 5}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
6 6 6 6 6 6`
	expect := []int{0, 0, 0, 0, 0, 0}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `9
8 4 2 6 3 9 5 7 8`
	expect := []int{0, 1, 2, 2, 4, 4, 4, 4, 5}
	runSample(t, s, expect)
}

func TestCommonGcdNeedsPrimePowers(t *testing.T) {
	s := `12
12 6 1 1 1 1 1 1 1 1 1 1`
	expect := []int{0, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	runSample(t, s, expect)
}
