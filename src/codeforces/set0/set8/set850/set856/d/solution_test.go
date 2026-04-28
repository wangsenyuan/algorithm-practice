package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `7 3
1 1 2 2 3 3
4 5 1
6 7 1
2 3 1`, 2)
}

func TestNoExtraEdges(t *testing.T) {
	runSample(t, `3 0
1 2`, 0)
}

func TestChooseBestOverlappingPath(t *testing.T) {
	runSample(t, `4 3
1 2 3
1 3 5
2 4 7
1 4 10`, 10)
}

func TestCanChooseDisjointPaths(t *testing.T) {
	runSample(t, `7 3
1 1 2 2 3 3
4 5 4
6 7 5
4 7 100`, 100)
}

func TestSmallAgainstBruteForce(t *testing.T) {
	for n := 3; n <= 9; n++ {
		for seed := 0; seed < 5; seed++ {
			parent := make([]int, n)
			children := make([][]int, n)
			treeEdge := make(map[[2]int]bool)
			for i := 1; i < n; i++ {
				parent[i] = (i*37 + seed) % i
				children[parent[i]] = append(children[parent[i]], i)
				treeEdge[[2]int{parent[i], i}] = true
			}

			var edges [][]int
			for u := 0; u < n && len(edges) < 10; u++ {
				for v := u + 1; v < n && len(edges) < 10; v++ {
					if treeEdge[[2]int{u, v}] {
						continue
					}
					w := (u*13+v*17+seed*19)%11 + 1
					edges = append(edges, []int{u, v, w})
				}
			}

			expect := bruteForce(parent, edges)
			res := solve(parent, children, edges)
			if res != expect {
				t.Fatalf("n=%d seed=%d expect %d, but got %d", n, seed, expect, res)
			}
		}
	}
}

func bruteForce(parent []int, edges [][]int) int {
	n := len(parent)
	depth := make([]int, n)
	for i := 1; i < n; i++ {
		depth[i] = depth[parent[i]] + 1
	}

	paths := make([][]int, len(edges))
	for i, e := range edges {
		u, v := e[0], e[1]
		var path []int
		for depth[u] > depth[v] {
			path = append(path, u)
			u = parent[u]
		}
		for depth[v] > depth[u] {
			path = append(path, v)
			v = parent[v]
		}
		for u != v {
			path = append(path, u, v)
			u = parent[u]
			v = parent[v]
		}
		path = append(path, u)
		paths[i] = path
	}

	var best int
	for mask := 0; mask < 1<<len(edges); mask++ {
		used := make([]bool, n)
		var sum int
		ok := true
		for i, e := range edges {
			if mask>>i&1 == 0 {
				continue
			}
			for _, u := range paths[i] {
				if used[u] {
					ok = false
					break
				}
				used[u] = true
			}
			if !ok {
				break
			}
			sum += e[2]
		}
		if ok {
			best = max(best, sum)
		}
	}
	return best
}
