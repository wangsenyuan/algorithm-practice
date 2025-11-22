package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	transformation := make([][]int, n-1)
	for i := range n - 1 {
		var x, k int
		fmt.Fscan(reader, &x, &k)
		transformation[i] = []int{x, k}
	}
	return solve(b, a, transformation)
}

const inf = 2e17

func solve(b []int, a []int, transformation [][]int) bool {
	n := len(b)

	g := NewGraph(n, n)

	for i, t := range transformation {
		j, k := t[0], t[1]
		g.AddEdge(j-1, i+1, k)
	}

	var dfs func(u int) (int, bool)
	dfs = func(u int) (int, bool) {
		bal := b[u] - a[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			tmp, ok := dfs(v)
			if !ok {
				return 0, false
			}
			if tmp >= 0 {
				bal += tmp
			} else {
				if bal < 0 && -tmp >= inf/g.val[i] {
					return 0, false
				}
				bal += tmp * g.val[i]
			}

			if bal < -inf {
				return 0, false
			}
		}

		return bal, true
	}

	res, ok := dfs(0)
	return ok && res >= 0
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
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
