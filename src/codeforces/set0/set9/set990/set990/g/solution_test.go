package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 3
1 2
2 3
	`
	expect := [][]int{
		{1, 4},
		{2, 1},
		{3, 1},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
1 2 4 8 16 32
1 6
6 3
3 4
4 2
6 5
	`
	expect := [][]int{
		{1, 6},
		{2, 5},
		{4, 6},
		{8, 1},
		{16, 2},
		{32, 1},
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
9 16 144 6
1 3
2 3
4 3
	`
	expect := [][]int{
		{1, 1},
		{2, 1},
		{3, 1},
		{6, 2},
		{9, 2},
		{16, 2},
		{144, 1},
	}
	runSample(t, s, expect)
}
