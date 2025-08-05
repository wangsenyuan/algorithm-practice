package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := process(reader)
	var buf bytes.Buffer
	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d\n", e))
	}
	fmt.Print(buf.String())
}

func process(reader *bufio.Reader) (n int, edges [][]int, res []int) {
	var m int
	fmt.Fscan(reader, &n, &m)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return n, edges, solve(n, edges)
}

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int) []int {
	f := make([]int, n)
	g := make([][]pair, n)
	for i, e := range edges {
		u, v, w := e[0]-1, e[1]-1, e[2]
		f[u] += w
		f[v] += w
		g[u] = append(g[u], pair{v, i})
		g[v] = append(g[v], pair{u, i})
	}

	for i := range f {
		f[i] /= 2
	}

	m := len(edges)

	ans := make([]int, m)
	for i := range m {
		ans[i] = -1
	}

	assign := func(u, v, i int) {
		e := edges[i]
		if e[0]-1 == u {
			ans[i] = 0
		} else {
			ans[i] = 1
		}
	}

	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++

	for tail < head {
		u := que[tail]
		tail++
		for _, cur := range g[u] {
			if ans[cur.second] == -1 {
				assign(u, cur.first, cur.second)
				v := cur.first
				f[v] -= edges[cur.second][2]
				if v != n-1 && f[v] == 0 {
					que[head] = v
					head++
				}
			}
		}
	}

	for i := range ans {
		if ans[i] == -1 {
			ans[i] = 1
		}
	}
	return ans
}
