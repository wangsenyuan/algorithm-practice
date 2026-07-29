package main

import (
	"bufio"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 5
2 3
1 3
1 2
1 2
1 2
`, []string{"YES", "YES", "NO", "YES", "NO"})
}

func TestRollbackAfterDeletingAnOddCycleEdge(t *testing.T) {
	runSample(t, `4 6
1 2
2 3
1 3
1 4
1 3
2 4
`, []string{"YES", "YES", "NO", "NO", "YES", "NO"})
}

func TestRandomSmallGraphsAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for it := 0; it < 1000; it++ {
		n := 2 + rng.Intn(5)
		q := 1 + rng.Intn(20)
		queries := make([][]int, q)
		for i := range queries {
			u := 1 + rng.Intn(n)
			v := 1 + rng.Intn(n)
			for u == v {
				v = 1 + rng.Intn(n)
			}
			if u > v {
				u, v = v, u
			}
			queries[i] = []int{u, v}
		}

		got := solve(n, queries)
		expect := bruteForce(n, queries)
		if !reflect.DeepEqual(got, expect) {
			t.Fatalf("n = %d, queries = %v, expect %v, but got %v", n, queries, expect, got)
		}
	}
}

func bruteForce(n int, queries [][]int) []string {
	edges := make([][]bool, n+1)
	for i := range edges {
		edges[i] = make([]bool, n+1)
	}
	ans := make([]string, len(queries))

	for i, query := range queries {
		u, v := query[0], query[1]
		edges[u][v] = !edges[u][v]
		edges[v][u] = edges[u][v]

		color := make([]int, n+1)
		for j := 1; j <= n; j++ {
			color[j] = -1
		}
		ok := true
		for start := 1; start <= n && ok; start++ {
			if color[start] >= 0 {
				continue
			}
			color[start] = 0
			queue := []int{start}
			for head := 0; head < len(queue) && ok; head++ {
				u = queue[head]
				for v := 1; v <= n; v++ {
					if !edges[u][v] {
						continue
					}
					if color[v] < 0 {
						color[v] = color[u] ^ 1
						queue = append(queue, v)
					} else if color[v] == color[u] {
						ok = false
						break
					}
				}
			}
		}
		if ok {
			ans[i] = "YES"
		} else {
			ans[i] = "NO"
		}
	}

	return ans
}
