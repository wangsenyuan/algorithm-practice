package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, v := range res {
		buf.WriteString(fmt.Sprintf("%d ", v))
	}
	buf.WriteByte('\n')
	os.Stdout.Write(buf.Bytes())
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
	n := readNum(reader)
	a := readNNums(reader, n)
	p := readNNums(reader, n)
	return solve(a, p)
}

func solve(a []int, p []int) []int {
	n := len(a)
	from := make([]int, n)
	dist := make([]int, n)
	for i := range n {
		from[i] = -1
		dist[i] = -1
	}

	deg := make([]int, n)
	for i := range n {
		if p[i] > 0 {
			deg[p[i]-1]++
		}
	}

	que := make([]int, n)

	var head, tail int

	for i := range n {
		if a[i] == 1 {
			que[head] = i
			head++
			dist[i] = 0
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		if p[u] > 0 {
			v := p[u] - 1
			if deg[v] == 1 {
				que[head] = v
				head++
				dist[v] = dist[u] + 1
				from[v] = u
			}
		}
	}
	x := slices.Max(dist)

	for u := range n {
		if dist[u] == x {
			var path []int
			for u != -1 {
				path = append(path, u+1)
				u = from[u]
			}
			return path
		}
	}

	return nil
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
