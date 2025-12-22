package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	checks := make([][]int, m)
	for i := 0; i < m; i++ {
		checks[i] = readNNums(reader, 3)
	}
	res := solve(n, checks)

	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	fmt.Println(buf.String())
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

func solve(n int, checks [][]int) []int {
	g := NewGraph(n+1, len(checks)+n)

	for _, c := range checks {
		u, v, w := c[0], c[1], c[2]
		g.AddEdge(v, u, w)
	}
	for i := 1; i <= n; i++ {
		g.AddEdge(0, i, 0)
	}

	neg := make([]int, n+1)
	neg[0]++
	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = 1 << 30
	}

	que := list.New()
	que.PushBack(0)
	in := make([]bool, n+1)
	in[0] = true
	for que.Len() > 0 {
		it := que.Front()
		que.Remove(it)
		u := it.Value.(int)
		in[u] = false

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w
				if !in[v] {
					in[v] = true
					que.PushBack(v)
					neg[v]++
					if neg[v] > n {
						// a negative loop
						return nil
					}
				}
			}
		}
	}

	return dist[1:]
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
