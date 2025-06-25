package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
	}
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (dominoes [][]int, res []string) {
	n := readNum(reader)
	dominoes = make([][]int, n)
	for i := 0; i < n; i++ {
		dominoes[i] = readNNums(reader, 2)
	}
	res = solve(dominoes)
	return
}

func solve(dominoes [][]int) []string {
	adj := make([][][]int, 7)
	for i := range 7 {
		adj[i] = make([][]int, 7)
	}

	// 如果 l == r 会怎么样呢？
	// 这部分是不是特殊处理一下？
	// 如果是不连通的，直接返回false
	self_loop := make([][]int, 7)
	deg := make([]int, 7)
	g := NewGraph(7, len(dominoes)*2)
	for i, cur := range dominoes {
		l, r := cur[0], cur[1]
		if l == r {
			self_loop[l] = append(self_loop[l], i+1)
			continue
		}
		adj[l][r] = append(adj[l][r], i+1)
		adj[r][l] = append(adj[r][l], -(i + 1))
		deg[l]++
		deg[r]++
		g.AddEdge(l, r, i)
		g.AddEdge(r, l, i)
	}

	var odd []int
	var root int
	var cnt int
	for i := 0; i <= 6; i++ {
		if deg[i] > 0 || len(self_loop[i]) > 0 {
			root = i
			cnt++
		}
		if deg[i]%2 == 1 {
			odd = append(odd, i)
		}
	}
	if len(odd) > 2 || len(odd) == 1 || !checkConnect(g, root, cnt) {
		return []string{"No solution"}
	}
	// 比如要有两个， 才能有欧拉回路
	if len(odd) == 2 {
		root = odd[0]
	}
	pos := make([]int, 7)
	for i := range 7 {
		pos[i] = g.nodes[i]
	}

	marked := make([]bool, len(dominoes))

	var path []int
	var dfs func(u int)
	dfs = func(u int) {
		for pos[u] > 0 {
			v := g.to[pos[u]]
			w := g.val[pos[u]]
			pos[u] = g.next[pos[u]]
			if !marked[w] {
				marked[w] = true
				dfs(v)
			}
		}
		path = append(path, u)
	}

	dfs(root)

	var ans []string

	for len(self_loop[path[0]]) > 0 {
		eid := self_loop[path[0]][0]
		self_loop[path[0]] = self_loop[path[0]][1:]
		ans = append(ans, fmt.Sprintf("%d +", eid))
	}

	clear(marked)

	for i := 0; i+1 < len(path); i++ {
		u := path[i]
		v := path[i+1]
		var eid int
		for len(adj[u][v]) > 0 {
			eid = adj[u][v][0]
			adj[u][v] = adj[u][v][1:]
			if !marked[abs(eid)-1] {
				marked[abs(eid)-1] = true
				break
			}
		}
		// panic when eid == 0
		if eid > 0 {
			ans = append(ans, fmt.Sprintf("%d +", eid))
		} else {
			ans = append(ans, fmt.Sprintf("%d -", -eid))
		}
		for len(self_loop[v]) > 0 {
			eid := self_loop[v][0]
			self_loop[v] = self_loop[v][1:]
			ans = append(ans, fmt.Sprintf("%d +", eid))
		}
	}

	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func checkConnect(g *Graph, r int, cnt int) bool {
	vis := make([]bool, 7)
	var dfs func(u int) int
	dfs = func(u int) int {
		if vis[u] {
			return 0
		}
		vis[u] = true
		res := 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			res += dfs(v)
		}
		return res
	}

	return dfs(r) == cnt
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
