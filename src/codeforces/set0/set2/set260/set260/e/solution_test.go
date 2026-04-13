package main

import (
	"bufio"
	"math/rand"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	a, points, ok, vertical, horizontal := drive(reader)
	if ok != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, ok)
	}
	if !expect {
		return
	}

	cnt := make([]int, 9)

	for _, cur := range points {
		x, y := float64(cur[0]), float64(cur[1])
		if x < vertical[0] {
			if y < horizontal[0] {
				cnt[0]++
			} else if y < horizontal[1] {
				cnt[1]++
			} else {
				cnt[2]++
			}
		} else if x < vertical[1] {
			if y < horizontal[0] {
				cnt[3]++
			} else if y < horizontal[1] {
				cnt[4]++
			} else {
				cnt[5]++
			}
		} else {
			if y < horizontal[0] {
				cnt[6]++
			} else if y < horizontal[1] {
				cnt[7]++
			} else {
				cnt[8]++
			}
		}
	}

	slices.Sort(cnt)
	slices.Sort(a)

	if !slices.Equal(cnt, a) {
		t.Fatalf("Sample expect %v, but got %v", a, cnt)
	}
}

func TestSample1(t *testing.T) {
	s := `9
1 1
1 2
1 3
2 1
2 2
2 3
3 1
3 2
3 3
1 1 1 1 1 1 1 1 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `15
4 4
-1 -3
1 5
3 -4
-4 4
-1 1
3 -3
-4 -5
-3 3
3 2
4 1
-4 2
-2 -5
-3 4
-1 4
2 1 2 1 2 1 3 2 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
-2 10
6 0
-16 -6
-4 13
-4 -2
-17 -10
9 15
18 16
-5 2
10 -5
2 1 1 1 1 1 1 1 1
`
	expect := false
	runSample(t, s, expect)
}

func TestNegativeYBoundaries(t *testing.T) {
	s := `9
-3 -3
-3 -2
-3 -1
-2 -3
-2 -2
-2 -1
-1 -3
-1 -2
-1 -1
1 1 1 1 1 1 1 1 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSolveMatchesSolve1OnRandomSmallCases(t *testing.T) {
	rng := rand.New(rand.NewSource(1))

	for tc := 0; tc < 40; tc++ {
		n := 9 + rng.Intn(8)
		used := make(map[[2]int]bool)
		points := make([][]int, 0, n)
		for len(points) < n {
			p := [2]int{rng.Intn(11) - 5, rng.Intn(11) - 5}
			if used[p] {
				continue
			}
			used[p] = true
			points = append(points, []int{p[0], p[1]})
		}

		cutsX := []int{-2, 0}
		cutsY := []int{-1, 2}
		cnt := make([]int, 9)
		for _, p := range points {
			x, y := p[0], p[1]
			id := 0
			if x >= cutsX[0] {
				id += 3
			}
			if x >= cutsX[1] {
				id += 3
			}
			if y >= cutsY[0] {
				id++
			}
			if y >= cutsY[1] {
				id++
			}
			cnt[id]++
		}
		if slices.Contains(cnt, 0) {
			tc--
			continue
		}

		a := slices.Clone(cnt)
		for i := len(a) - 1; i > 0; i-- {
			j := rng.Intn(i + 1)
			a[i], a[j] = a[j], a[i]
		}

		ok1, _, _ := solve1(slices.Clone(a), clonePoints(points))
		ok2, _, _ := solve(slices.Clone(a), clonePoints(points))
		if ok1 != ok2 {
			t.Fatalf("mismatch on case %d: solve1=%v solve=%v a=%v points=%v", tc, ok1, ok2, a, points)
		}
		if !ok1 {
			continue
		}
	}
}

func clonePoints(points [][]int) [][]int {
	res := make([][]int, len(points))
	for i, p := range points {
		res[i] = slices.Clone(p)
	}
	return res
}
