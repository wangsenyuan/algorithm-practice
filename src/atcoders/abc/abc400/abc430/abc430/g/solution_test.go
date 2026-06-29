package main

import (
	"bufio"
	"math/bits"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][2]int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 7
1 1 2 10
1 2 4 20
3 1 3
2 1 2 20
1 2 3 10
3 1 2
3 1 4
`, [][2]int{
		{2, 1},
		{1, 2},
		{2, 1},
	})
}

func TestSmallAgainstBruteForce(t *testing.T) {
	queries := [][]int{
		{1, 1, 4, 1},
		{1, 2, 5, 2},
		{3, 1, 5},
		{2, 3, 4, 1},
		{3, 1, 5},
		{1, 3, 3, 1},
		{2, 1, 5, 2},
		{3, 2, 4},
	}
	expect := brute(5, queries)
	res := solve(5, queries)
	if !slices.Equal(res, expect) {
		t.Fatalf("expect %v, but got %v", expect, res)
	}
}

func TestRandomSmallAgainstBruteForce(t *testing.T) {
	seed := uint32(1)
	next := func(n int) int {
		seed ^= seed << 13
		seed ^= seed >> 17
		seed ^= seed << 5
		return int(seed % uint32(n))
	}

	for tc := 0; tc < 100; tc++ {
		n := 1 + next(20)
		q := 1 + next(200)
		queries := make([][]int, 0, q)
		for i := 0; i < q; i++ {
			tp := 1 + next(3)
			l := 1 + next(n)
			r := 1 + next(n)
			if l > r {
				l, r = r, l
			}
			if tp == 3 {
				queries = append(queries, []int{tp, l, r})
			} else {
				x := 1 + next(6)
				queries = append(queries, []int{tp, l, r, x})
			}
		}
		expect := brute(n, queries)
		res := solve(n, queries)
		if !slices.Equal(res, expect) {
			t.Fatalf("n=%d queries=%v expect %v, but got %v", n, queries, expect, res)
		}
	}
}

func brute(n int, queries [][]int) [][2]int {
	sets := make([]uint64, n)
	var res [][2]int
	for _, cur := range queries {
		l, r := cur[1]-1, cur[2]-1
		if cur[0] == 1 {
			x := uint64(1) << (cur[3] - 1)
			for i := l; i <= r; i++ {
				sets[i] |= x
			}
		} else if cur[0] == 2 {
			x := uint64(1) << (cur[3] - 1)
			for i := l; i <= r; i++ {
				sets[i] &^= x
			}
		} else {
			best, cnt := 0, 0
			for i := l; i <= r; i++ {
				v := bits.OnesCount64(sets[i])
				if v > best {
					best = v
					cnt = 1
				} else if v == best {
					cnt++
				}
			}
			res = append(res, [2]int{best, cnt})
		}
	}
	return res
}
