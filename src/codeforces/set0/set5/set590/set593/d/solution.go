package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 3)
		for j := range 3 {
			fmt.Fscan(reader, &edges[i][j])
		}
	}
	queries := make([][]int, m)
	for i := range m {
		var t int
		fmt.Fscan(reader, &t)
		if t == 1 {
			queries[i] = make([]int, 4)
		} else {
			queries[i] = make([]int, 3)
		}
		queries[i][0] = t
		for j := 1; j < len(queries[i]); j++ {
			fmt.Fscan(reader, &queries[i][j])
		}
	}
	return solve(n, edges, queries)
}

type edge struct {
	to int
	id int
}

func solve(n int, edges [][]int, queries [][]int) []int {

	g := make([][]int, n)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], i)
		g[v] = append(g[v], i)
	}

	dep := make([]int, n)

	fa := make([]int, n)
	for i := range n {
		fa[i] = i
	}

	find := func(u int) int {
		v := u
		for fa[v] != v {
			v = fa[v]
		}
		for fa[u] != v {
			fa[u], u = v, fa[u]
		}
		return v
	}

	path := make([]edge, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {

		for _, eid := range g[u] {
			x, y, z := edges[eid][0]-1, edges[eid][1]-1, edges[eid][2]
			v := x ^ y ^ u
			if p != v {
				dep[v] = dep[u] + 1
				if z == 1 {
					fa[v] = find(u)
				}
				path[v] = edge{u, eid}
				dfs(u, v)
			}
		}
	}

	dfs(0, 0)

	query := func(u int, v int, num int) int {
		u = find(u)
		v = find(v)

		for num > 0 && u != v {
			if dep[u] < dep[v] {
				u, v = v, u
			}
			// dep[u] >= dep[v]
			p := path[u]
			num /= edges[p.id][2]
			u = find(p.to)
		}

		return num
	}

	update := func(eid int, z int) {
		x, y := edges[eid][0]-1, edges[eid][1]-1
		if dep[x] > dep[y] {
			x, y = y, x
		}
		// x -> y
		if z > 1 || x == 0 {
			edges[eid][2] = z
		} else {
			// z == 1
			fa[y] = find(x)
		}
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == 2 {
			eid, z := cur[1], cur[2]
			update(eid-1, z)
		} else {
			u, v, num := cur[1]-1, cur[2]-1, cur[3]
			ans = append(ans, query(u, v, num))
		}
	}

	return ans
}
