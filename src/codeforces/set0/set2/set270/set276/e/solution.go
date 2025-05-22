package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n, q := readTwoNums(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 2)
	}
	queries := make([][]int, q)
	for i := range q {
		s, _ := reader.ReadBytes('\n')
		var tp int
		pos := readInt(s, 0, &tp) + 1
		if tp == 0 {
			queries[i] = make([]int, 4)
		} else {
			queries[i] = make([]int, 2)
		}
		queries[i][0] = tp
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}

	return solve(n, edges, queries)
}

func solve(n int, edges [][]int, queries [][]int) []int {
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	D := make([]int, n)
	fa := make([]int, n)
	sz := make([]int, n)

	var dfs func(r int, u int)
	dfs = func(r int, u int) {
		fa[u] = r
		sz[r] = max(sz[r], D[u])
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == 0 || D[v] > 0 {
				continue
			}
			D[v] = D[u] + 1
			dfs(r, v)
		}
	}

	// 这样子保证空间复杂度为O(n)
	sets := make([]SegTree, n)
	sets[0] = make(SegTree, 2*n)
	for i := g.nodes[0]; i > 0; i = g.next[i] {
		r := g.to[i]
		D[r] = 1
		dfs(r, r)
		// 把0也算进去， 比较好计算
		sz[r]++
		sets[r] = make(SegTree, 2*sz[r])
	}

	var ans []int

	for _, cur := range queries {
		if cur[0] == 0 {
			v, x, d := cur[1], cur[2], cur[3]
			v--
			if v == 0 {
				// 对区间[0..d]增加
				sets[0].Update(0, d+1, x)
			} else {
				// v != 0
				l := max(1, D[v]-d)
				r := min(sz[fa[v]]-1, D[v]+d)
				sets[fa[v]].Update(l, r+1, x)
				if D[v] <= d {
					d0 := d - D[v]
					sets[0].Update(0, d0+1, x)
					if d0 > 0 {
						// 要把这条线上重复的部分给取消掉
						sets[fa[v]].Update(1, min(sz[fa[v]], d0+1), -x)
					}
				}
			}
		} else {
			v := cur[1] - 1
			tmp := sets[0].Get(D[v])
			if v != 0 {
				tmp += sets[fa[v]].Get(D[v])
			}
			ans = append(ans, tmp)
		}
	}

	return ans
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	next := make([]int, e+3)
	to := make([]int, e+3)
	var cur int
	return &Graph{nodes, next, to, cur}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}

type SegTree []int

func (st SegTree) Update(l int, r int, v int) {
	n := len(st) / 2
	l += n
	r += n
	for l < r {
		if l&1 == 1 {
			st[l] += v
			l++
		}
		if r&1 == 1 {
			r--
			st[r] += v
		}
		l >>= 1
		r >>= 1
	}
}

func (st SegTree) Get(p int) int {
	n := len(st) / 2
	p += n
	var res int
	for p > 0 {
		res += st[p]
		p /= 2
	}
	return res
}
