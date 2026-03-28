package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

const inf = 1 << 60

type Pair struct {
	dist int
	idx  int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	roads := make([][]int, m)
	for i := 0; i < m; i++ {
		roads[i] = make([]int, 3)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1], &roads[i][2])
	}
	var g1, g2, s1, s2 int
	fmt.Fscan(reader, &g1, &g2, &s1, &s2)
	return solve(n, roads, g1, g2, s1, s2)
}

func solve(n int, roads [][]int, g1 int, g2 int, s1 int, s2 int) int64 {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	for _, road := range roads {
		u, v, w := road[0]-1, road[1]-1, road[2]
		dist[u][v] = w
		dist[v][u] = w
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	times := make([][]int, n)
	lo := make([]Pair, n)
	hi := make([]Pair, n)

	for i := 0; i < n; i++ {
		var cur []int
		for j := range n {
			if i != j {
				cur = append(cur, dist[i][j])
			}
		}
		slices.Sort(cur)
		cur = slices.Compact(cur)
		times[i] = cur
		lo[i] = Pair{cur[0], i}
		hi[i] = Pair{cur[len(cur)-1], i}
	}

	xs := make([]Pair, 0, n)
	for i := range n {
		xs = append(xs, lo[i])
	}
	xs = sortAndUniquePairs(xs)

	var ans int64

	for _, x := range xs {
		nextPos := make([]Pair, n)
		canGold := make([]bool, n)
		xPos := -1
		ys := make([]Pair, 0, n)
		for i := range n {
			nextPos[i] = nextAfter(times[i], i, x)
			canGold[i] = !lessPair(x, lo[i])
			if equalPair(lo[i], x) {
				xPos = i
			}
			if nextPos[i].idx >= 0 {
				ys = append(ys, nextPos[i])
			}
		}
		ys = sortAndUniquePairs(ys)

		for _, y := range ys {
			if !lessPair(x, y) {
				continue
			}
			ans += countWays(n, hi, nextPos, canGold, xPos, y, g1, g2, s1, s2)
		}
	}

	return ans
}

func countWays(n int, hi []Pair, nextPos []Pair, canGold []bool, xPos int, y Pair, g1 int, g2 int, s1 int, s2 int) int64 {
	width := (s2 + 1) * 4
	size := (g2 + 1) * width
	dp := make([]int64, size)
	ndp := make([]int64, size)
	dp[0] = 1

	for i := range n {
		clear(ndp)

		canSilver := nextPos[i].idx >= 0 && !lessPair(y, nextPos[i])
		canBronze := lessPair(y, hi[i])
		isY := equalPair(nextPos[i], y)

		maxG := min(i, g2)
		maxS := min(i, s2)
		for g := 0; g <= maxG; g++ {
			baseG := g * width
			limitS := min(maxS, i-g)
			for s := 0; s <= limitS; s++ {
				base := baseG + s*4
				for state := range 4 {
					fx := state >> 1
					fy := state & 1
					if canGold[i] && g < g2 {
						nfx := fx
						if i == xPos {
							nfx = 1
						}
						ndp[(g+1)*width+s*4+(nfx<<1)+fy] += dp[base+state]
					}
					if canSilver && s < s2 {
						nfy := fy
						if isY {
							nfy = 1
						}
						ndp[baseG+(s+1)*4+(fx<<1)+nfy] += dp[base+state]
					}
					if canBronze {
						ndp[base+state] += dp[base+state]
					}
				}
			}
		}
		dp, ndp = ndp, dp
	}

	var res int64
	for g := g1; g <= g2; g++ {
		baseG := g * width
		for s := s1; s <= s2; s++ {
			res += dp[baseG+s*4+3]
		}
	}
	return res
}

func nextAfter(arr []int, idx int, x Pair) Pair {
	pos := sort.SearchInts(arr, x.dist)
	if pos < len(arr) && arr[pos] == x.dist && idx > x.idx {
		return Pair{x.dist, idx}
	}
	pos = sort.Search(len(arr), func(j int) bool {
		return arr[j] > x.dist
	})
	if pos == len(arr) {
		return Pair{-1, -1}
	}
	return Pair{arr[pos], idx}
}

func lessPair(a Pair, b Pair) bool {
	return a.dist < b.dist || a.dist == b.dist && a.idx < b.idx
}

func equalPair(a Pair, b Pair) bool {
	return a.dist == b.dist && a.idx == b.idx
}

func sortAndUniquePairs(arr []Pair) []Pair {
	slices.SortFunc(arr, func(a, b Pair) int {
		return cmp.Or(a.dist-b.dist, a.idx-b.idx)
	})
	return slices.Compact(arr)
}
