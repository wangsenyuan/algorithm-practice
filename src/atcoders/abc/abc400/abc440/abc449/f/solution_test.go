package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 2 2 3
1 3
2 4
3 1
`
	runSample(t, s, 2)
}

func TestSample2(t *testing.T) {
	s := `4 4 3 2 2
2 2
3 4
`
	runSample(t, s, 0)
}

func TestSample3(t *testing.T) {
	s := `449 449 3 14 0
`
	runSample(t, s, 194892)
}

func TestSample4(t *testing.T) {
	s := `31 9 5 7 10
14 8
8 4
18 8
12 1
8 5
9 6
18 1
14 7
5 6
26 7
`
	runSample(t, s, 12)
}

func TestSolveAgainstBruteForce(t *testing.T) {
	cases := []struct {
		H      int64
		W      int64
		h      int64
		w      int64
		blacks [][]int64
	}{
		{1, 1, 1, 1, nil},
		{1, 1, 1, 1, [][]int64{{1, 1}}},
		{3, 4, 2, 2, [][]int64{{1, 3}, {2, 4}, {3, 1}}},
		{4, 4, 3, 2, [][]int64{{2, 2}, {3, 4}}},
		{5, 5, 2, 3, [][]int64{{1, 1}, {2, 3}, {5, 5}}},
		{6, 5, 4, 2, [][]int64{{2, 1}, {3, 2}, {3, 5}, {6, 4}}},
	}
	for _, cur := range cases {
		expect := brute(cur.H, cur.W, cur.h, cur.w, cur.blacks)
		res := solve(cur.H, cur.W, cur.h, cur.w, cur.blacks)
		if res != expect {
			t.Fatalf("Grid %dx%d size %dx%d blacks %v expect %d, but got %d",
				cur.H, cur.W, cur.h, cur.w, cur.blacks, expect, res)
		}
	}
}

func brute(H int64, W int64, h int64, w int64, blacks [][]int64) int64 {
	bad := make(map[[2]int64]bool)
	for _, p := range blacks {
		bad[[2]int64{p[0], p[1]}] = true
	}
	var res int64
	for r := int64(1); r+h-1 <= H; r++ {
		for c := int64(1); c+w-1 <= W; c++ {
			ok := true
			for i := int64(0); i < h && ok; i++ {
				for j := int64(0); j < w; j++ {
					if bad[[2]int64{r + i, c + j}] {
						ok = false
						break
					}
				}
			}
			if ok {
				res++
			}
		}
	}
	return res
}
