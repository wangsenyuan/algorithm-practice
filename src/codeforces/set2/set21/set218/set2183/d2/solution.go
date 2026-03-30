package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tc int
	fmt.Fscan(in, &tc)
	for range tc {
		_, res := drive(in)
		fmt.Fprintln(out, len(res))
		for _, cur := range res {
			fmt.Fprint(out, len(cur))
			for _, x := range cur {
				fmt.Fprint(out, " ", x)
			}
			fmt.Fprintln(out)
		}
	}
}

func drive(reader *bufio.Reader) (edges [][]int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	edges = make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	adj := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0]-1, edge[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	parent, levels, k := buildLevels(adj)
	color := make([]int, n)
	for i := range color {
		color[i] = -1
	}
	color[0] = 0

	groups := make([][]int, k)
	inUsed := make([]bool, k)
	pos := make([]int, k)
	forbidden := make([]int, n)

	for depth := 1; depth < len(levels); depth++ {
		assignLevel(levels[depth], parent, color, k, groups, inUsed, pos, forbidden)
	}

	res := make([][]int, k)
	for i := 0; i < n; i++ {
		res[color[i]] = append(res[color[i]], i+1)
	}

	return res
}

func buildLevels(adj [][]int) ([]int, [][]int, int) {
	n := len(adj)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = -1
	}

	que := make([]int, n)
	head, tail := 1, 0
	que[0] = 0

	var levels [][]int
	k := 1

	for tail < head {
		end := head
		level := make([]int, 0, end-tail)
		for tail < end {
			u := que[tail]
			tail++
			level = append(level, u)

			children := 0
			for _, v := range adj[u] {
				if v == parent[u] {
					continue
				}
				parent[v] = u
				que[head] = v
				head++
				children++
			}
			k = max(k, children+1)
		}
		k = max(k, len(level))
		levels = append(levels, level)
	}

	return parent, levels, k
}

func assignLevel(nodes []int, parent []int, color []int, k int, groups [][]int, inUsed []bool, pos []int, forbidden []int) {
	used := make([]int, 0, len(nodes))
	for _, u := range nodes {
		ban := color[parent[u]]
		if !inUsed[ban] {
			inUsed[ban] = true
			used = append(used, ban)
			groups[ban] = groups[ban][:0]
		}
		groups[ban] = append(groups[ban], u)
	}

	need := len(nodes)
	if need < k {
		need++
	}

	free := make([]int, 0, need)
	for _, c := range used {
		free = append(free, c)
	}
	for c := 0; len(free) < need && c < k; c++ {
		if !inUsed[c] {
			free = append(free, c)
		}
	}

	for i, c := range free {
		pos[c] = i
	}

	assigned := make([]int, 0, len(nodes))
	popColor := func(c int) {
		i := pos[c]
		last := len(free) - 1
		x := free[last]
		free[i] = x
		pos[x] = i
		free = free[:last]
	}
	takeColor := func(ban int) int {
		last := len(free) - 1
		if last < 0 {
			return -1
		}
		if free[last] != ban {
			c := free[last]
			popColor(c)
			return c
		}
		if last == 0 {
			return -1
		}
		c := free[last-1]
		popColor(c)
		return c
	}

	for _, ban := range used {
		for _, u := range groups[ban] {
			c := takeColor(ban)
			if c >= 0 {
				color[u] = c
				forbidden[u] = ban
				assigned = append(assigned, u)
				continue
			}

			popColor(ban)
			for _, v := range assigned {
				if forbidden[v] == ban {
					continue
				}
				color[u] = color[v]
				color[v] = ban
				forbidden[u] = ban
				assigned = append(assigned, u)
				break
			}
		}
	}

	for _, ban := range used {
		inUsed[ban] = false
	}
}
