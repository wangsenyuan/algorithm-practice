package main

import (
	"bufio"
	"math/rand"
	"reflect"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !reflect.DeepEqual(res, expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 3
0 0 4 1
1 2
1 3
1 4
1 4
2 3
2 4
`, []int64{3, 6, 6})
}

func TestSample2(t *testing.T) {
	runSample(t, `4 3
0 4 1 2
1 3
1 4
2 4
1 2
2 3
2 4
`, []int64{6, 10, 3})
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3
3 2 4 4
1 2
2 4
3 4
1 2
1 3
2 3
`, []int64{2, 5, 4})
}

func bruteSolve(a []int, edges [][]int, queries [][]int) []int64 {
	n := len(a)
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	ans := make([]int64, len(queries))
	for qi, query := range queries {
		src, dst := query[0]-1, query[1]-1
		parent := make([]int, n)
		for i := range parent {
			parent[i] = -1
		}
		parent[src] = src
		queue := []int{src}
		for front := 0; front < len(queue); front++ {
			u := queue[front]
			for _, v := range adj[u] {
				if parent[v] < 0 {
					parent[v] = u
					queue = append(queue, v)
				}
			}
		}

		var path []int
		for u := dst; u != src; u = parent[u] {
			path = append(path, u)
		}
		path = append(path, src)

		for l := range path {
			var xor, sum int
			for r := l; r < len(path); r++ {
				xor ^= a[path[r]]
				sum += a[path[r]]
				if xor == sum {
					ans[qi]++
				}
			}
		}
	}
	return ans
}

func checkAgainstBrute(t *testing.T, a []int, edges [][]int, queries [][]int) {
	t.Helper()
	expect := bruteSolve(a, edges, queries)
	got := solve(a, edges, queries)
	if !reflect.DeepEqual(got, expect) {
		t.Fatalf("a=%v edges=%v queries=%v: expect %v, got %v", a, edges, queries, expect, got)
	}
}

func TestFocusedRegressions(t *testing.T) {
	tests := []struct {
		name    string
		a       []int
		edges   [][]int
		queries [][]int
	}{
		{
			name:    "all-zero path",
			a:       []int{0, 0},
			edges:   [][]int{{1, 2}},
			queries: [][]int{{1, 2}},
		},
		{
			name:    "same bit on opposite LCA branches",
			a:       []int{0, 1, 1},
			edges:   [][]int{{1, 2}, {1, 3}},
			queries: [][]int{{2, 3}},
		},
		{
			name:    "earlier same-arm collision survives a zero",
			a:       []int{0, 1, 1, 0},
			edges:   [][]int{{1, 2}, {2, 3}, {3, 4}},
			queries: [][]int{{2, 4}},
		},
		{
			name:    "DFS state does not leak between sibling subtrees",
			a:       []int{0, 1, 1, 0},
			edges:   [][]int{{1, 2}, {1, 3}, {2, 4}},
			queries: [][]int{{3, 4}, {1, 3}},
		},
		{
			name: "long zero runs are aggregated",
			a:    []int{0, 0, 0, 1, 0, 0, 2, 0, 0, 0, 4, 0},
			edges: [][]int{
				{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7},
				{7, 8}, {8, 9}, {9, 10}, {10, 11}, {11, 12},
			},
			queries: [][]int{{1, 12}, {3, 10}, {5, 12}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkAgainstBrute(t, test.a, test.edges, test.queries)
		})
	}
}

func TestRandomAgainstBrute(t *testing.T) {
	rng := rand.New(rand.NewSource(2236))
	for it := 0; it < 300; it++ {
		n := 2 + rng.Intn(9)
		a := make([]int, n)
		for i := range a {
			a[i] = rng.Intn(1 << 5)
		}

		edges := make([][]int, 0, n-1)
		for v := 2; v <= n; v++ {
			edges = append(edges, []int{1 + rng.Intn(v-1), v})
		}

		var queries [][]int
		for x := 1; x <= n; x++ {
			for y := x + 1; y <= n; y++ {
				queries = append(queries, []int{x, y})
			}
		}
		checkAgainstBrute(t, a, edges, queries)
	}
}

func TestMaxZeroChain(t *testing.T) {
	const n = 100000
	a := make([]int, n)
	edges := make([][]int, 0, n-1)
	for v := 2; v <= n; v++ {
		edges = append(edges, []int{v - 1, v})
	}

	got := solve(a, edges, [][]int{{1, n}})
	expect := int64(n) * int64(n+1) / 2
	if len(got) != 1 || got[0] != expect {
		t.Fatalf("expect %d, got %v", expect, got)
	}
}

func TestAllTwentyBitsAcrossBranches(t *testing.T) {
	const bitsCount = 20
	a := make([]int, 1+2*bitsCount)
	edges := make([][]int, 0, 2*bitsCount)

	parent := 1
	for bit := 0; bit < bitsCount; bit++ {
		vertex := 2 + bit
		a[vertex-1] = 1 << bit
		edges = append(edges, []int{parent, vertex})
		parent = vertex
	}
	leftLeaf := parent

	parent = 1
	for bit := 0; bit < bitsCount; bit++ {
		vertex := 2 + bitsCount + bit
		a[vertex-1] = 1 << bit
		edges = append(edges, []int{parent, vertex})
		parent = vertex
	}
	rightLeaf := parent

	checkAgainstBrute(t, a, edges, [][]int{{leftLeaf, rightLeaf}})
}
