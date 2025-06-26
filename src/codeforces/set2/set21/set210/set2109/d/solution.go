package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
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
	n, m, k := readThreeNums(reader)
	a := readNNums(reader, k)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	return solve(a, n, edges)
}

const inf = 1 << 30

func solve(a []int, n int, edges [][]int) string {
	var sum int
	// mx[1]是最小的奇数
	mx := make([]int, 2)
	mx[1] = inf
	for _, x := range a {
		sum += x
		if x&1 == 1 {
			mx[1] = min(mx[1], x)
		}
	}

	if sum&1 == 0 {
		mx[0] = sum
		if mx[1] == inf {
			// 没有奇数，全部是偶数
			mx[1] = -1
		} else {
			// 偶数减去最小的奇数 = 奇数
			mx[1] = sum - mx[1]
		}
	} else {
		// 最大的偶数
		mx[0] = sum - mx[1]
		mx[1] = sum
	}

	g := NewGraph(n, 2*len(edges))

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	dist := make([][]int, n)

	for i := range dist {
		dist[i] = make([]int, 2)
		for j := range 2 {
			dist[i][j] = -1
		}
	}

	dist[0][0] = 0

	que := make([]int, 2*n)
	var head, tail int
	que[head] = 0
	head++

	for tail < head {
		u, d := que[tail]/2, que[tail]%2
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if dist[v][d^1] == -1 {
				dist[v][d^1] = dist[u][d] + 1
				que[head] = v*2 + (d ^ 1)
				head++
			}
		}
	}

	buf := make([]byte, n)
	for i := range n {
		if i == 0 {
			buf[i] = '1'
			continue
		}
		buf[i] = '0'
		if dist[i][0] != -1 && dist[i][0] <= mx[0] {
			buf[i] = '1'
		} else if dist[i][1] != -1 && dist[i][1] <= mx[1] {
			buf[i] = '1'
		}
	}
	return string(buf)
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
