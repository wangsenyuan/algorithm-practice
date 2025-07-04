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

const inf = 1 << 60

func solve(a []int, edges [][]int) int {
	n := len(a)
	g := NewGraph(n, 2*n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	// sz[u] = 将u处理后, 剩余的apple的数量
	sz := make([]int, n)
	// val[u] = 将u处理后，如果要再次处理，需要消耗val[u]的倍数
	val := make([]int, n)
	var res int
	var dfs func(p int, u int) bool
	dfs = func(p int, u int) bool {
		val[u] = 1
		x := inf
		var cnt int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				cnt++
				if !dfs(u, v) {
					return false
				}
				val[u] = lcm(val[u], val[v])
				x = min(x, sz[v])
				sz[u] += sz[v]
			}
		}
		if cnt == 0 {
			sz[u] = a[u]
			return true
		}
		// not a leaf
		if x < val[u] {
			//必须保留val[u]的倍数（或者是0）
			return false
		}

		x = x / val[u] * val[u]
		res += sz[u] - x*cnt
		sz[u] = x * cnt
		val[u] *= cnt
		return true
	}

	if !dfs(0, 0) {
		res = 0
		// 全部删除掉
		for _, v := range a {
			res += v
		}
	}

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

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
