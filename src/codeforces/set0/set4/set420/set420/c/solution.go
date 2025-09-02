package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, p int
	fmt.Fscan(reader, &n, &p)
	statements := make([][]int, n)
	for i := range n {
		statements[i] = make([]int, 2)
		fmt.Fscan(reader, &statements[i][0], &statements[i][1])
	}
	return solve(n, statements, p)
}

type pair struct {
	first  int
	second int
}

func solve(n int, statements [][]int, p int) int {
	adj := make([][]int, n)

	freq := make(map[pair]int)

	for _, cur := range statements {
		x, y := cur[0]-1, cur[1]-1
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
		x, y = min(x, y), max(x, y)
		freq[pair{x, y}]++
	}
	arr := make([]pair, n)
	deg := make([]int, n)
	for i := range n {
		deg[i] = len(adj[i])
		arr[i] = pair{len(adj[i]), i}
		sort.Ints(adj[i])
		adj[i] = slices.Compact(adj[i])
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	pos := make([]int, n)
	for i, cur := range arr {
		pos[cur.second] = i
	}

	var res int
	// deg[x] + deg[y] - deg[x & y] >= p
	for l1, r := n-1, 0; r < n; r++ {
		for l1 >= 0 && arr[l1].first+arr[r].first >= p {
			l1--
		}

		// l2 == -1 or arr[l2].first + arr[r].first == p
		// l2...i的部分肯定满足条件
		// 但是l1...i中的部分，不一定满足条件
		l := max(l1, r)
		res += n - (l + 1)
		u := arr[r].second
		for _, v := range adj[u] {
			x, y := min(u, v), max(u, v)
			if pos[v] > l && deg[x]+deg[y]-freq[pair{x, y}] < p {
				// v 可能出现两次
				res--
			}
		}
	}

	return res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
