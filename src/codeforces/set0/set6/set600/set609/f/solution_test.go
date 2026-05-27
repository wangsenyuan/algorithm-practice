package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 6
10 2
15 0
6 1
0 1
110 10
1 1
6 0
15 10
14 100
12 2
`
	expect := [][]int{
		{3, 114},
		{1, 10},
		{1, 1},
		{1, 2},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `1 2
10 2
20 2
12 1
`
	expect := [][]int{
		{1, 3},
	}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 10
33 2
922 34
480 105
844 5
739 39
325 20
999 88
462 104
225 5
93 4
13 15
323 9
152 20
785 11
512 4
859 8
327 14
818 9
794 13
99 20
`
	expect := [][]int{
		{0, 2},
		{0, 34},
		{0, 105},
		{0, 5},
		{0, 39},
		{1, 34},
		{0, 88},
		{1, 108},
		{0, 5},
		{0, 4},
	}
	runSample(t, s, expect)
}

func TestEarlierGrowingFrogKeepsPriority(t *testing.T) {
	s := `6 9
5 8
28 5
14 5
25 5
13 7
23 2
25 7
33 3
31 4
31 8
32 5
29 7
22 8
29 7
14 5
`
	expect := [][]int{
		{0, 8},
		{1, 8},
		{0, 5},
		{0, 5},
		{2, 20},
		{6, 40},
	}
	runSample(t, s, expect)
}
