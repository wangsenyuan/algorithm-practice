package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	E := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		E[i] = readNNums(reader, 3)
	}
	Q := make([]int, m)
	for i := 0; i < m; i++ {
		Q[i] = readNum(reader)
	}
	res := solve(n, E, Q)
	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("AYE\n")
		} else {
			buf.WriteString("NAY\n")
		}
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(n int, E [][]int, Q []int) []bool {
	g := NewGraph(n, 2*n)

	for _, e := range E {
		u, v, w := e[0]-1, e[1]-1, e[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}

	sz := make([]int, n)
	mp := make([]int, n)
	vis := make([]bool, n)

	root := -1

	var centroid func(p int, u int, tot int)

	centroid = func(p int, u int, tot int) {
		sz[u] = 1
		mp[u] = 0
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p == v || vis[v] {
				continue
			}
			centroid(u, v, tot)
			sz[u] += sz[v]
			mp[u] = max(mp[u], sz[v])
		}
		mp[u] = max(mp[u], tot-sz[u])
		if root < 0 || mp[u] < mp[root] {
			root = u
		}
	}

	ans := make([]bool, len(Q))

	var dfs2 func(u int)

	var dfs3 func(p int, u int, dist int, r int)

	var it int

	arr := make([]Node, n)

	dfs3 = func(p int, u int, dist int, r int) {
		arr[it] = Node{dist, u, r}
		it++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v && !vis[v] {
				dfs3(u, v, dist+g.val[i], r)
			}
		}
	}

	calc := func(u int) {
		arr[0] = Node{0, u, u}
		it = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				dfs3(u, v, g.val[i], v)
			}
		}

		sort.Slice(arr[:it], func(i, j int) bool {
			return arr[i].dist < arr[j].dist
		})
		for i, k := range Q {
			if ans[i] {
				continue
			}
			l, r := 0, it-1
			for l < r {
				a, b := arr[l], arr[r]
				if a.dist+b.dist > k {
					r--
					continue
				}
				if a.dist+b.dist < k {
					l++
					continue
				}
				// a.first + b.first == k
				if a.from != b.from {
					ans[i] = true
					break
				}
				if arr[r-1].dist == arr[r].dist {
					r--
				} else {
					l++
				}
			}
		}
	}

	dfs2 = func(u int) {
		vis[u] = true
		calc(u)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !vis[v] {
				root = -1
				centroid(-1, v, sz[v])
				dfs2(root)
			}
		}
	}

	centroid(-1, 0, n)
	dfs2(root)

	return ans
}

type Node struct {
	dist int
	id   int
	from int
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
	e++
	next := make([]int, e)
	to := make([]int, e)
	val := make([]int, e)
	return &Graph{nodes, next, to, val, 0}
}

func (g *Graph) AddEdge(u, v, w int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
	g.val[g.cur] = w
}
