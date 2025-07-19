package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(res[0], res[1])
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

func process(reader *bufio.Reader) ([]int, [][]int) {
	n := readNum(reader)
	lambs := make([][]int, n)
	for i := 0; i < n; i++ {
		lambs[i] = readNNums(reader, 2)
	}
	return solve(lambs), lambs
}

func solve(lambs [][]int) []int {
	n := len(lambs)

	var sum int
	for _, t := range lambs {
		sum += t[1]
	}

	if sum%3 != 0 {
		return nil
	}

	g := NewGraph(n, n)

	deg := make([]int, n)

	var root int

	for u, lamb := range lambs {
		p := lamb[0]
		if p > 0 {
			// g.AddEdge(p-1, u)
			g.AddEdge(u, p-1)
			deg[p-1]++
		} else {
			root = u
		}
	}

	val := make([]int, n)
	que := make([]int, n)
	var head, tail int
	for u := range n {
		if deg[u] == 0 {
			que[head] = u
			head++
		}
	}

	var res []int

	for tail < head {
		u := que[tail]
		tail++
		val[u] += lambs[u][1]
		if u != root && val[u]*3 == sum && len(res) < 2 {
			res = append(res, u+1)
			val[u] = 0
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			val[v] += val[u]
			deg[v]--
			if deg[v] == 0 {
				que[head] = v
				head++
			}
		}
	}

	if len(res) < 2 {
		return nil
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
