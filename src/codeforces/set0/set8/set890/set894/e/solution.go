package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	roads := make([][]int32, m)
	for i := range m {
		roads[i] = make([]int32, 3)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1], &roads[i][2])
	}
	var s int
	fmt.Fscan(reader, &s)
	return solve(n, roads, s)
}

func solve(n int, roads [][]int32, s int) int {
	s--

	// loop需要特殊处理
	g := NewGraph(n, len(roads))

	var mw int

	for _, road := range roads {
		u, v, w := road[0]-1, road[1]-1, road[2]
		mw = max(mw, int(w))
		if u != v {
			g.AddEdge(u, v, w)
		}
	}

	dsc := make([]int32, n)
	low := make([]int32, n)
	var timer int32
	stack := make([]int32, n)
	var top int
	marked := make([]bool, n)

	belong := make([]int32, n)
	var gid int

	var dfs func(u int)
	dfs = func(u int) {
		timer++
		dsc[u] = timer
		low[u] = timer
		stack[top] = int32(u)
		top++
		marked[u] = true

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := int(g.to[i])
			if dsc[v] == 0 {
				dfs(v)
				low[u] = min(low[u], low[v])
			} else if marked[v] {
				low[u] = min(low[u], dsc[v])
			}
		}

		if low[u] == dsc[u] {
			gid++
			for top > 0 {
				v := int(stack[top-1])
				top--
				marked[v] = false
				belong[v] = int32(gid)
				if v == u {
					break
				}
			}
		}
	}

	dfs(s)

	fp := make([]int, gid)

	updateAnswer := func(i int, w int) {
		j := sort.Search(w+1, func(j int) bool {
			return j*(j+1)/2 > w
		})
		j--
		sum := j * (j + 1) * (j + 2) / 6
		fp[i] += (j+1)*w - sum
	}

	tr := NewGraph(gid, len(roads))

	deg := make([]int32, gid)

	for _, e := range roads {
		u, v, w := e[0]-1, e[1]-1, e[2]
		if belong[u] == 0 {
			// 这条边不会被访问到(感觉只需要检查belong[u]就可以了)
			continue
		}
		if belong[u] == belong[v] {
			updateAnswer(int(belong[u])-1, int(w))
		} else {
			tr.AddEdge(belong[u]-1, belong[v]-1, w)
			deg[belong[v]-1]++
		}
	}

	ans := make([]int, gid)
	que := make([]int32, gid)
	var head, tail int
	que[head] = belong[s] - 1
	head++

	var best int

	for tail < head {
		u := que[tail]
		tail++
		ans[u] += fp[u]
		best = max(best, ans[u])
		for i := tr.nodes[u]; i > 0; i = tr.next[i] {
			v := tr.to[i]
			w := tr.val[i]
			ans[v] = max(ans[v], ans[u]+int(w))
			deg[v]--

			if deg[v] == 0 {
				que[head] = int32(v)
				head++
			}
		}
	}

	return best
}

type Graph struct {
	nodes []int32
	next  []int32
	to    []int32
	val   []int32
	cur   int32
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int32, n)
	e++
	next := make([]int32, e)
	to := make([]int32, e)
	val := make([]int32, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int32) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
