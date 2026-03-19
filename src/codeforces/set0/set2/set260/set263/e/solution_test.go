package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4 2
1 2 3 4
1 1 1 1
2 2 2 2
4 3 2 1
`
	runSample(t, s, []int{3, 2})
}

func TestSample2(t *testing.T) {
	s := `5 7 3
8 2 3 4 2 3 3
3 4 6 2 3 4 6
8 7 6 8 4 5 7
1 2 3 2 1 3 2
4 5 3 2 1 2 1
`
	runSample(t, s, []int{3, 3})
}

func bruteSolve(k int, a [][]int) []int {
	n := len(a)
	m := len(a[0])
	best := -1
	ans := []int{k, k}
	for x := k - 1; x <= n-k; x++ {
		for y := k - 1; y <= m-k; y++ {
			var cur int
			for i := range n {
				for j := range m {
					dist := abs(i-x) + abs(j-y)
					if dist < k {
						cur += (k - dist) * a[i][j]
					}
				}
			}
			if cur > best {
				best = cur
				ans[0] = x + 1
				ans[1] = y + 1
			}
		}
	}
	return ans
}

func TestSmallBruteforceCases(t *testing.T) {
	cases := []struct {
		k int
		a [][]int
	}{
		{
			1,
			[][]int{
				{2, 2},
				{2, 3},
			},
		},
		{
			2,
			[][]int{
				{1, 0, 2, 1},
				{3, 4, 1, 2},
				{2, 1, 5, 0},
				{0, 3, 2, 1},
			},
		},
		{
			2,
			[][]int{
				{5, 1, 0},
				{2, 4, 3},
				{1, 0, 2},
				{3, 1, 1},
				{0, 2, 4},
			},
		},
		{
			3,
			[][]int{
				{0, 1, 2, 0, 1},
				{2, 3, 1, 4, 0},
				{1, 5, 2, 3, 1},
				{0, 2, 4, 1, 2},
				{1, 0, 3, 2, 4},
			},
		},
	}

	for _, tc := range cases {
		got := solve(tc.k, tc.a)
		expect := bruteSolve(tc.k, tc.a)
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("solve(%d, %v) expect %v, but got %v", tc.k, tc.a, expect, got)
		}
	}
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
