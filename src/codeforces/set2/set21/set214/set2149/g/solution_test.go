package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := drive(reader)
	for i, u := range expect {
		v := res[i]
		if !slices.Equal(u, v) {
			t.Fatalf("Sample expect %v, but got %v, at %d", u, v, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `1 1
5
1 1`
	expect := [][]int{{5}}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
1 1 2 3
1 4
2 3`
	expect := [][]int{{1}, {1, 2}}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 3
7 7 7 8 8 9
1 6
2 5
4 6`
	expect := [][]int{{7}, {7, 8}, {8}}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `8 2
4 4 4 5 5 5 6 6
1 8
3 6`
	expect := [][]int{{4, 5}, {5}}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 5
1 2 3 3 3 4 4 4 4 5
1 10
1 5
4 9
6 9
7 10`
	expect := [][]int{{4}, {3}, {4}, {4}, {4}}
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `20 20
10 7 5 3 8 10 1 10 7 6 4 2 9 4 2 10 3 6 6 6
2 2
10 15
4 19
6 18
10 13
2 3
5 6
14 15
16 17
7 11
11 11
2 14
7 17
6 15
2 20
4 11
4 8
17 20
5 8
1 17`
	expect := [][]int{
		{7},
		{-1},
		{-1},
		{-1},
		{-1},
		{5, 7},
		{8, 10},
		{2, 4},
		{3, 10},
		{-1},
		{4},
		{-1},
		{-1},
		{-1},
		{-1},
		{-1},
		{10},
		{6},
		{10},
		{-1},
	}
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `20 2
10 7 5 3 8 10 1 10 7 6 4 2 9 4 2 10 3 6 6 6
5 6
16 17
`
	expect := [][]int{
		{8, 10},
		{3, 10},
	}
	runSample(t, s, expect)
}

func TestMergeKeepsTwoDistinctCandidatesForThirdThreshold(t *testing.T) {
	got := merge([]pair{{1, 1}}, []pair{{2, 1}}, 2)
	expect := []pair{{1, 1}, {2, 1}}
	if !slices.Equal(got, expect) {
		t.Fatalf("merge got %v, want %v", got, expect)
	}
}
