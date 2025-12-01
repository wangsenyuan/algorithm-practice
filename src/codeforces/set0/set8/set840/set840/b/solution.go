package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	_, _, _, ok, res := drive(reader)
	if !ok {
		fmt.Println("-1")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(res))
	for _, x := range res {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) (n int, d []int, edges [][]int, ok bool, res []int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	d = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &d[i])
	}
	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	ok, res = solve(n, slices.Clone(d), edges)
	return
}

func solve(n int, d []int, edges [][]int) (bool, []int) {
	cnt := make([]int, 2)

	for i := range n {
		if d[i] < 0 {
			cnt[0]++
		} else if d[i] > 0 {
			cnt[1]++
		}
	}
	if cnt[0] == 0 && cnt[1]%2 == 1 {
		return false, nil
	}
	if cnt[1]%2 == 1 {
		// 需要将一个变成1， 其他的变成0
		for i := range n {
			if d[i] < 0 {
				d[i] = 1
				break
			}
		}
	}
	for i := range n {
		if d[i] < 0 {
			d[i] = 0
		}
	}

	g := NewGraph(n, 2*len(edges))

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	marked := make([]bool, len(edges))
	vis := make([]bool, n)

	var dfs func(u int)

	dfs = func(u int) {
		vis[u] = true
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs(v)
				if d[v] == 1 {
					marked[g.val[i]] = true
					d[v] ^= 1
					d[u] ^= 1
				}
			}
		}
	}

	dfs(0)

	var res []int
	for i := range len(edges) {
		if marked[i] {
			res = append(res, i+1)
		}
	}

	return true, res
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	val   []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	val := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, val, cur}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
