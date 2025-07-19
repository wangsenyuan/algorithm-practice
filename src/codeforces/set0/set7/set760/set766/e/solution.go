package main

import (
	"bufio"
	"fmt"
	"os"
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	a := readNNums(reader, n)
	edges := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(a, edges)
}

const H = 20

func solve(a []int, edges [][]int) int {
	n := len(a)
	g := NewGraph(n, 2*n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}
	sz := make([]int, n)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		sz[u] = 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs(u, v)
				sz[u] += sz[v]
			}
		}
	}

	dfs(-1, 0)

	s1 := make([]int, n)

	var dfs1 func(p int, u int, d int)
	dfs1 = func(p int, u int, d int) {
		s1[u] = (a[u] >> d) & 1
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs1(u, v, d)
				if (a[u]>>d)&1 == 1 {
					s1[u] += sz[v] - s1[v]
				} else {
					s1[u] += s1[v]
				}
			}
		}
	}

	var res int

	var dfs2 func(p int, u int, d int, x int)

	dfs2 = func(p int, u int, d int, x int) {
		if p >= 0 {
			if (a[u]>>d)&1 == 1 {
				s1[u] += n - sz[u] - x
			} else {
				s1[u] += x
			}
		}
		res += s1[u] * (1 << d)
		// sz[u] = n

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				nx := s1[u]
				// 要把v的贡献去掉
				if (a[u]>>d)&1 == 1 {
					nx -= sz[v] - s1[v]
				} else {
					nx -= s1[v]
				}
				dfs2(u, v, d, nx)
			}
		}
	}

	for d := range H {
		dfs1(-1, 0, d)
		dfs2(-1, 0, d, 0)
	}

	var sum int
	for _, x := range a {
		sum += x
	}
	res -= sum
	res /= 2
	res += sum

	return res
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
