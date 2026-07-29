package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, line := range drive(reader) {
		fmt.Fprintln(writer, line)
	}
}

func drive(reader *bufio.Reader) []string {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) []string {
	q := len(queries)
	segments := make([][]edge, 4*q)
	active := make(map[edge]int)

	for i, query := range queries {
		e := edge{query[0], query[1]}
		if start, ok := active[e]; ok {
			addInterval(segments, 1, 0, q, start, i, e)
			delete(active, e)
		} else {
			active[e] = i
		}
	}

	for e, start := range active {
		addInterval(segments, 1, 0, q, start, q, e)
	}

	dsu := newRollbackDSU(n)
	ans := make([]string, q)
	var dfs func(int, int, int)
	dfs = func(node, left, right int) {
		snapshot := len(dsu.history)
		for _, e := range segments[node] {
			dsu.add(e.u, e.v)
		}

		if right-left == 1 {
			if dsu.bad == 0 {
				ans[left] = "YES"
			} else {
				ans[left] = "NO"
			}
		} else {
			mid := (left + right) / 2
			dfs(node*2, left, mid)
			dfs(node*2+1, mid, right)
		}

		dsu.rollback(snapshot)
	}

	dfs(1, 0, q)
	return ans
}

type edge struct {
	u int
	v int
}

func addInterval(segments [][]edge, node, left, right, start, end int, e edge) {
	if start <= left && right <= end {
		segments[node] = append(segments[node], e)
		return
	}
	mid := (left + right) / 2
	if start < mid {
		addInterval(segments, node*2, left, mid, start, end, e)
	}
	if mid < end {
		addInterval(segments, node*2+1, mid, right, start, end, e)
	}
}

type rollbackDSU struct {
	parent  []int
	size    []int
	parity  []int
	bad     int
	history []change
}

type change struct {
	root       int
	parent     int
	parentSize int
	rootParity int
	bad        int
	merged     bool
}

func newRollbackDSU(n int) *rollbackDSU {
	dsu := &rollbackDSU{
		parent: make([]int, n+1),
		size:   make([]int, n+1),
		parity: make([]int, n+1),
	}
	for i := 1; i <= n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (dsu *rollbackDSU) find(x int) (int, int) {
	var parity int
	for dsu.parent[x] != x {
		parity ^= dsu.parity[x]
		x = dsu.parent[x]
	}
	return x, parity
}

func (dsu *rollbackDSU) add(u, v int) {
	ru, pu := dsu.find(u)
	rv, pv := dsu.find(v)
	if ru == rv {
		if pu == pv {
			dsu.history = append(dsu.history, change{bad: dsu.bad})
			dsu.bad++
		}
		return
	}

	if dsu.size[ru] > dsu.size[rv] {
		ru, rv = rv, ru
	}
	// size[ru] < size[rv]
	dsu.history = append(dsu.history, change{
		root:       ru,
		parent:     rv,
		parentSize: dsu.size[rv],
		rootParity: dsu.parity[ru],
		merged:     true,
	})
	dsu.parent[ru] = rv
	dsu.parity[ru] = pu ^ pv ^ 1
	dsu.size[rv] += dsu.size[ru]
}

func (dsu *rollbackDSU) rollback(snapshot int) {
	for len(dsu.history) > snapshot {
		i := len(dsu.history) - 1
		last := dsu.history[i]
		dsu.history = dsu.history[:i]
		if last.merged {
			dsu.parent[last.root] = last.root
			dsu.parity[last.root] = last.rootParity
			dsu.size[last.parent] = last.parentSize
		} else {
			dsu.bad = last.bad
		}
	}
}
