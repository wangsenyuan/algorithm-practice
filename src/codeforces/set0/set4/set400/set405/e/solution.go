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
	res, _, _ := process(reader)
	if len(res) == 0 {
		fmt.Println("No solution")
		return
	}
	var buf bytes.Buffer
	for _, e := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", e[0], e[1], e[2]))
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (res [][]int, n int, edges [][]int) {
	n, m := readTwoNums(reader)
	edges = make([][]int, m)
	for i := range m {
		edges[i] = readNNums(reader, 2)
	}
	res = solve(n, edges)
	return
}

func solve(n int, edges [][]int) [][]int {
	m := len(edges)
	if m%2 == 1 {
		return nil
	}
	g := NewGraph(n, 2*m)

	for i, e := range edges {
		u, v := e[0], e[1]
		u--
		v--
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)

	}

	marked := make([]bool, m)

	var ans [][]int

	var dfs func(u int) int

	dfs = func(u int) int {
		var adj []int
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if !marked[w] {
				marked[w] = true
				adj = append(adj, v)
			}
		}
		var unpaired []int
		for _, v := range adj {
			w := dfs(v)
			if w < 0 {
				unpaired = append(unpaired, v)
			} else {
				ans = append(ans, []int{u + 1, v + 1, w + 1})
			}
		}
		for len(unpaired) > 1 {
			v := unpaired[0]
			w := unpaired[1]
			ans = append(ans, []int{v + 1, u + 1, w + 1})
			unpaired = unpaired[2:]
		}
		if len(unpaired) == 1 {
			return unpaired[0]
		}
		return -1
	}

	dfs(0)

	return ans
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
