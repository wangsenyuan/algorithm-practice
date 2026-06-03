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
		t.Errorf("Sample %s, expect %v, but got %v", s, expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `1
1 2`
	expect := []int{1, 2}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
1 3 4 2`
	expect := []int{1, 3, 2, 4}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
2 3 4 1`
	expect := []int{1, 4, 2, 3}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3
8 3 4 2 1 5 7 6`
	expect := []int{1, 5, 6, 7, 2, 4, 3, 8}
	runSample(t, s, expect)
}
