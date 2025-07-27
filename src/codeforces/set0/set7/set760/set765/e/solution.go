package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := make([]int, n)
	track := make([]int, n)

	que := make([]int, n)

	bfs := func(s int) {
		for i := range n {
			dist[i] = -1
			track[i] = -1
		}
		dist[s] = 0
		var head, tail int
		que[head] = s
		head++
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[v] < 0 {
					dist[v] = dist[u] + 1
					track[v] = u
					que[head] = v
					head++
				}
			}
		}
	}

	bfs(0)
	var first int
	for i := range n {
		if dist[i] > dist[first] {
			first = i
		}
	}
	bfs(first)
	var second int
	for i := range n {
		if dist[i] > dist[second] {
			second = i
		}
	}

	var dfs func(p int, u int) int

	dfs = func(p int, u int) int {
		var d int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v {
				continue
			}
			tmp := dfs(u, v)
			// 它的每个子树的深度必须一样，才能被合并
			if tmp < 0 || d > 0 && d != tmp {
				return -1
			}
			d = tmp
		}
		return d + 1
	}

	type pair struct {
		first  int
		second int
	}
	next := -1
	var branches []pair
	for u := second; u >= 0; u = track[u] {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == next || v == track[u] {
				continue
			}
			d := dfs(u, v)
			if d < 0 {
				return -1
			}
			if d == dist[u] || d == dist[second]-dist[u] {
				continue
			}
			branches = append(branches, pair{d, u})
		}
		next = u
	}

	slices.SortFunc(branches, func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	})

	branches = slices.Compact(branches)

	if len(branches) > 1 {
		return -1
	}

	if len(branches) == 0 {
		return get(dist[second])
	}

	r := branches[0].second
	if dist[r]*2 == dist[second] {
		// 在中间位置
		tot := dist[second]/2 + branches[0].first
		return get(tot)
	}
	// 不在中间
	return -1
}

func get(x int) int {
	for x%2 == 0 {
		x /= 2
	}
	return x
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
