package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) string {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	ok := solve(n, edges)
	if ok {
		return "YES"
	}
	return "NO"
}

func solve(n int, edges [][]int) bool {
	g := NewGraph(n, len(edges)+1)

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v, i)
	}

	dfn := make([]int, n)
	for i := range n {
		dfn[i] = -1
	}
	low := make([]int, n)
	stack := make([]int, n)
	var top int
	vis := make([]bool, n)
	var comps [][]int

	var timer int
	var dfs func(u int)
	dfs = func(u int) {
		dfn[u] = timer
		low[u] = timer
		timer++
		stack[top] = u
		top++
		vis[u] = true

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if vis[v] {
				low[u] = min(low[u], dfn[v])
			} else if dfn[v] < 0 {
				dfs(v)
				low[u] = min(low[u], low[v])
			}
			// else 这部分已经被处理了，且肯定没有访问到u
		}

		if low[u] == dfn[u] {
			var comp []int
			for top > 0 {
				v := stack[top-1]
				top--
				vis[v] = false
				comp = append(comp, v)
				if v == u {
					break
				}
			}
			comps = append(comps, comp)
		}
	}

	for u := range n {
		if dfn[u] < 0 {
			dfs(u)
		}
	}

	it := -1
	for i, cur := range comps {
		if len(cur) == 1 {
			continue
		}
		if it >= 0 {
			// 2个强连通分量
			return false
		}
		it = i
	}
	if it < 0 {
		return true
	}

	// 有一个强连通分量
	marked := make([]bool, n)
	for _, u := range comps[it] {
		marked[u] = true
	}

	in := make([]int, n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		if marked[u] && marked[v] {
			in[v]++
		}
	}

	que := make([]int, n)
	deg := make([]int, n)

	bfs := func(s int, rem int) bool {
		copy(deg, in)
		var head, tail int
		que[head] = s
		head++
		// deg[s] == 1
		deg[s]--

		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				w := g.val[i]
				if w != rem && marked[v] {
					deg[v]--
					if deg[v] == 0 {
						que[head] = v
						head++
					}
				}
			}
		}

		return head == len(comps[it])
	}

	for i, e := range edges {
		u, v := e[0]-1, e[1]-1
		if marked[u] && marked[v] && in[v] == 1 {
			if bfs(v, i) {
				return true
			}
		}
	}

	return false
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
