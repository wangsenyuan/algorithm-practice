package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.Abs(res-expect) > 1e-9 {
		t.Fatalf("Sample expect %.12f, but got %.12f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
1 2
0 1
1 0
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 2
1 3
1 0
0 2
0 3
`, 2)
}

func TestSample3(t *testing.T) {
	runSample(t, `7
1 2
1 3
2 4
2 5
3 6
3 7
1 1
1 1
1 1
1 1
1 1
1 1
1 1
`, 4.04081632653)
}

func TestSmallAgainstExactDFS(t *testing.T) {
	cases := []struct {
		n     int
		edges [][]int
		x     []float64
		y     []float64
	}{
		{4, [][]int{{1, 2}, {2, 3}, {2, 4}}, []float64{1, 3, 2, 4}, []float64{5, 0, 1, 2}},
		{5, [][]int{{1, 2}, {1, 3}, {3, 4}, {3, 5}}, []float64{0, 2, 1, 3, 4}, []float64{1, 1, 0, 2, 3}},
		{6, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}, []float64{6, 5, 4, 3, 2, 1}, []float64{1, 2, 3, 4, 5, 6}},
	}
	for _, cur := range cases {
		x := normalize(cur.x)
		y := normalize(cur.y)
		res := solve(cur.n, cur.edges, x, y)
		expect := exactExpected(cur.n, cur.edges, x, y)
		if math.Abs(res-expect) > 1e-9 {
			t.Fatalf("sample expect %.12f, but got %.12f", expect, res)
		}
	}
}

func normalize(arr []float64) []float64 {
	res := make([]float64, len(arr))
	var sum float64
	for _, num := range arr {
		sum += num
	}
	for i, num := range arr {
		res[i] = num / sum
	}
	return res
}

func exactExpected(n int, edges [][]int, x []float64, y []float64) float64 {
	g := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	var res float64
	for s := 0; s < n; s++ {
		for t := 0; t < n; t++ {
			res += x[s] * y[t] * exactDFS(g, s, -1, t)
		}
	}
	return res
}

func exactDFS(g [][]int, u int, p int, target int) float64 {
	if u == target {
		return 0
	}
	children := make([]int, 0, len(g[u]))
	for _, v := range g[u] {
		if v != p {
			children = append(children, v)
		}
	}
	var sum float64
	var dfsPerm func([]int, int)
	dfsPerm = func(arr []int, pos int) {
		if pos == len(arr) {
			var cur float64
			for _, v := range arr {
				cur++
				if containsTarget(g, v, u, target) {
					cur += exactDFS(g, v, u, target)
					break
				}
				cur += float64(fullSize(g, v, u)*2 - 1)
			}
			sum += cur
			return
		}
		for i := pos; i < len(arr); i++ {
			arr[pos], arr[i] = arr[i], arr[pos]
			dfsPerm(arr, pos+1)
			arr[pos], arr[i] = arr[i], arr[pos]
		}
	}
	dfsPerm(children, 0)
	return sum / float64(fact(len(children)))
}

func containsTarget(g [][]int, u int, p int, target int) bool {
	if u == target {
		return true
	}
	for _, v := range g[u] {
		if v != p && containsTarget(g, v, u, target) {
			return true
		}
	}
	return false
}

func fullSize(g [][]int, u int, p int) int {
	res := 1
	for _, v := range g[u] {
		if v != p {
			res += fullSize(g, v, u)
		}
	}
	return res
}

func fact(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}
