package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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
	n, k := readTwoNums(reader)

	readEdge := func() []int {
		return readNNums(reader, 2)
	}

	return solve(n, k, readEdge)
}

func solve(n int, k int, readEdge func() []int) []int {
	g := NewGraph(n+1, 2*n)

	for i := 0; i < n-1; i++ {
		e := readEdge()
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	h := bits.Len(uint(n + 1))

	fa := make([][]int, n+1)
	dep := make([]int, n+1)
	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := int(g.to[i])
			if p != v {
				dep[v] = dep[u] + 1
				dfs(u, v)
			}
		}
	}

	dfs(n, n)

	marked := make([]bool, n+1)
	marked[n] = true
	cnt := 1
	for it := n - 1; cnt+k < n; it-- {
		if marked[it] {
			continue
		}
		u := it
		for i := h - 1; i >= 0; i-- {
			if !marked[fa[u][i]] {
				u = fa[u][i]
			}
		}
		if cnt+dep[it]-dep[u]+1+k <= n {
			cnt += dep[it] - dep[u] + 1
			u = it
			for !marked[u] {
				marked[u] = true
				u = fa[u][0]
			}
		}
	}

	res := make([]int, 0, k)
	for i := 1; i <= n; i++ {
		if !marked[i] {
			res = append(res, i)
		}
	}
	return res
}

type Graph struct {
	nodes []int32
	next  []int32
	to    []int32
	cur   int32
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int32, n)
	e++
	next := make([]int32, e)
	to := make([]int32, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = int32(v)
}
