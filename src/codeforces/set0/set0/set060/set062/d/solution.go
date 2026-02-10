package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if len(res) == 0 {
		fmt.Println("No solution")
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	old := make([]int, m+1)
	for i := range m + 1 {
		fmt.Fscan(reader, &old[i])
	}
	return solve(n, old)
}

func solve(n int, old []int) []int {
	m := len(old) - 1
	start := old[0] - 1
	base := newGraph(n)
	for i := 0; i < m; i++ {
		u, v := old[i]-1, old[i+1]-1
		base.add(u, v)
	}

	for i := m - 1; i >= 1; i-- {
		g := base.clone()
		for k := 0; k <= i-2; k++ {
			u, v := old[k]-1, old[k+1]-1
			g.remove(u, v)
		}

		cur := old[i-1] - 1
		oldNext := old[i] - 1
		for cand := oldNext + 1; cand < n; cand++ {
			if !g.has(cur, cand) {
				continue
			}
			g2 := g.clone()
			g2.remove(cur, cand)
			if !feasible(g2, cand, start) {
				continue
			}
			suf, ok := buildLexSmallest(g2, cand, start)
			if !ok {
				continue
			}
			ans := make([]int, 0, m+1)
			for k := 0; k < i; k++ {
				ans = append(ans, old[k])
			}
			for _, x := range suf {
				ans = append(ans, x+1)
			}
			if len(ans) == m+1 {
				return ans
			}
		}
	}

	return nil
}

type graph struct {
	n     int
	adj   [][]bool
	deg   []int
	edges int
}

func newGraph(n int) *graph {
	adj := make([][]bool, n)
	for i := range n {
		adj[i] = make([]bool, n)
	}
	return &graph{
		n:   n,
		adj: adj,
		deg: make([]int, n),
	}
}

func (g *graph) clone() *graph {
	ng := newGraph(g.n)
	ng.edges = g.edges
	copy(ng.deg, g.deg)
	for i := 0; i < g.n; i++ {
		copy(ng.adj[i], g.adj[i])
	}
	return ng
}

func (g *graph) has(u int, v int) bool {
	return g.adj[u][v]
}

func (g *graph) add(u int, v int) {
	if g.adj[u][v] {
		return
	}
	g.adj[u][v] = true
	g.adj[v][u] = true
	g.deg[u]++
	g.deg[v]++
	g.edges++
}

func (g *graph) remove(u int, v int) {
	if !g.adj[u][v] {
		return
	}
	g.adj[u][v] = false
	g.adj[v][u] = false
	g.deg[u]--
	g.deg[v]--
	g.edges--
}

func feasible(g *graph, start int, end int) bool {
	if g.edges == 0 {
		return start == end
	}
	if g.deg[start] == 0 || g.deg[end] == 0 {
		return false
	}

	oddCnt := 0
	for i := 0; i < g.n; i++ {
		if g.deg[i]%2 == 1 {
			oddCnt++
		}
	}
	if start == end {
		if oddCnt != 0 {
			return false
		}
	} else {
		if oddCnt != 2 || g.deg[start]%2 == 0 || g.deg[end]%2 == 0 {
			return false
		}
	}

	vis := make([]bool, g.n)
	que := make([]int, g.n)
	head, tail := 0, 0
	que[tail] = start
	tail++
	vis[start] = true
	for head < tail {
		u := que[head]
		head++
		for v := 0; v < g.n; v++ {
			if g.adj[u][v] && !vis[v] {
				vis[v] = true
				que[tail] = v
				tail++
			}
		}
	}

	for i := 0; i < g.n; i++ {
		if g.deg[i] > 0 && !vis[i] {
			return false
		}
	}
	return vis[end]
}

func buildLexSmallest(g *graph, start int, end int) ([]int, bool) {
	path := []int{start}
	cur := start
	for g.edges > 0 {
		pick := -1
		for nxt := 0; nxt < g.n; nxt++ {
			if !g.adj[cur][nxt] {
				continue
			}
			g.remove(cur, nxt)
			if feasible(g, nxt, end) {
				pick = nxt
				path = append(path, nxt)
				cur = nxt
				break
			}
			g.add(cur, nxt)
		}
		if pick < 0 {
			return nil, false
		}
	}
	return path, cur == end
}
