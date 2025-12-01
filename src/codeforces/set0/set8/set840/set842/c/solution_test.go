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
	s := `2
6 2
1 2`
	expect := []int{6, 6}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
6 2 3
1 2
1 3`
	expect := []int{6, 6, 6}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `1
10`
	expect := []int{10}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
2 3 4 5 6 7 8 9 10 11
1 2
2 3
3 4
4 5
5 6
6 7
4 8
8 9
9 10`
	expect := []int{2, 3, 2, 1, 1, 1, 1, 1, 1, 1}
	runSample(t, s, expect)
}
