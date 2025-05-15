package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 2)
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	m := len(edges)
	g := NewGraph(n, m)

	deg := make([]int, n)
	for i, edge := range edges {
		u, v := edge[0], edge[1]
		u--
		v--
		g.AddEdge(u, v, i)
		deg[v]++
	}

	que := make([]int, n)

	check := func(k int) bool {
		if k < n-1 {
			// 至少要n-1条边
			return false
		}
		clear(deg)
		for i := range k {
			deg[edges[i][1]-1]++
		}
		var head, tail int
		for i := 0; i < n; i++ {
			if deg[i] == 0 {
				que[head] = i
				head++
			}
		}

		for tail < head {
			if head-tail != 1 {
				return false
			}
			u := que[tail]
			tail++

			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]

				if g.val[i] >= k {
					continue
				}

				deg[v]--
				if deg[v] == 0 {
					que[head] = v
					head++
				}
			}
		}

		return head == n
	}
	if !check(m) {
		return -1
	}

	return sort.Search(m, check)
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
