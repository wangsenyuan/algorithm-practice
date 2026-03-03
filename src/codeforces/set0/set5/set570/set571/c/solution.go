package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	if len(res) == 0 {
		fmt.Fprintln(writer, "NO")
	} else {
		fmt.Fprintln(writer, "YES")
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) (clauses [][]int, res string) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	clauses = make([][]int, n)
	for i := range n {
		var k int
		fmt.Fscan(reader, &k)
		clauses[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &clauses[i][j])
		}
	}
	res = solve(m, clauses)
	return
}

type edge struct {
	to    int
	varId int
	val   int
}

func solve(m int, clauses [][]int) string {
	n := len(clauses)
	g := make([][]int, 2*m+1)
	for i, cur := range clauses {
		for _, v := range cur {
			if v > 0 {
				g[v] = append(g[v], i)
			} else {
				g[m-v] = append(g[m-v], i)
			}
		}
	}
	// 同一个variable出现两次
	res := make([]int, m+1)
	for i := range m + 1 {
		res[i] = -1
	}

	adj := make([][]edge, n)

	add := func(u int, v int, varId int, val int) {
		adj[u] = append(adj[u], edge{v, varId, val})
	}

	done := make([]bool, n)

	for i := 1; i <= m; i++ {
		if len(g[i]) == 0 {
			// pure negative (or unused): set x_i = 0 to satisfy neg clauses
			res[i] = 0
			for _, id := range g[m+i] {
				done[id] = true
			}
		} else if len(g[m+i]) == 0 {
			// pure positive (or unused): set x_i = 1 to satisfy pos clauses
			res[i] = 1
			for _, id := range g[i] {
				done[id] = true
			}
		} else {
			// variable appears in both signs — add edges between the two clauses
			// val means: if we traverse this edge (orient toward destination),
			// assign x_i = val to satisfy the destination clause
			add(g[i][0], g[m+i][0], i, 0) // posClause→negClause: x_i=0 satisfies negClause (¬x_i true)
			add(g[m+i][0], g[i][0], i, 1) // negClause→posClause: x_i=1 satisfies posClause (x_i true)
		}
	}

	used := make([]bool, n)

	var dfs func(u int)
	dfs = func(u int) {
		used[u] = true
		for _, it := range adj[u] {
			v := it.to
			if !used[v] {
				res[it.varId] = it.val
				dfs(v)
			}
		}
	}

	for i := range n {
		if !used[i] && done[i] {
			dfs(i)
		}
	}

	color := make([]int, n)

	var findCycle func(u int, lastEdge int) bool

	trace := make([]edge, n)
	for i := range n {
		trace[i].to = -1
	}

	var cycle []int

	findCycle = func(u int, lastEdge int) bool {
		color[u] = 1
		for _, cur := range adj[u] {
			if cur.varId == lastEdge {
				continue
			}
			v := cur.to
			if color[v] == 1 {
				// Found back edge u->v; set res for closing edge and trace path
				res[cur.varId] = cur.val
				if len(cycle) == 0 {
					w := u
					for {
						done[w] = true
						used[w] = true
						cycle = append(cycle, w)
						// 因为是cycle，肯定会到v
						res[trace[w].varId] = trace[w].val
						if w == v {
							break
						}
						w = trace[w].to
					}
				}
				return true
			}
			if color[v] == 0 {
				trace[v] = edge{u, cur.varId, cur.val}
				if findCycle(v, cur.varId) {
					return true
				}
			}
		}
		color[u] = 2
		return false
	}

	for i := range n {
		if used[i] {
			continue
		}
		cycle = cycle[:0]
		found := findCycle(i, -1)
		if !found {
			return ""
		}
		for _, v := range cycle {
			dfs(v)
		}
	}

	buf := make([]byte, m)
	for i := 1; i <= m; i++ {
		if res[i] == -1 {
			res[i] = 0
		}
		buf[i-1] = byte(res[i] + '0')
	}
	return string(buf)
}
