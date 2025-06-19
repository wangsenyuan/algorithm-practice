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
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	edges := make([][]int, m)
	for i := 0; i < m; i++ {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, edges)
}

const inf = 1 << 60

func solve(n int, edges [][]int) int {
	slices.SortFunc(edges, func(a, b []int) int {
		return a[2] - b[2]
	})

	set := NewDSU(n)
	g := NewGraph(n, len(edges)*2)

	marked := make([]bool, n)

	que := make([]int, n)

	var head, tail int

	que[head] = 0
	head++

	marked[0] = true

	res := inf
	best := inf

	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		u--
		v--
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)

		if !marked[u] && marked[v] {
			best = min(best, w)
			marked[u] = true
			que[head] = u
			head++
		}

		if !marked[v] && marked[u] {
			best = min(best, w)
			marked[v] = true
			que[head] = v
			head++
		}

		set.Union(u, v)
		if set.Find(0) == set.Find(n-1) {
			for tail < head {
				u := que[tail]
				tail++
				for i := g.nodes[u]; i > 0; i = g.next[i] {
					v := g.to[i]
					w := g.val[i]
					best = min(best, w)
					if !marked[v] {
						marked[v] = true
						que[head] = v
						head++
					}
				}
			}
			res = min(res, w+best)
		}
	}

	return res
}

type DSU struct {
	arr []int
	cnt []int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt}
}

func (this *DSU) Find(x int) int {
	if this.arr[x] != x {
		this.arr[x] = this.Find(this.arr[x])
	}
	return this.arr[x]
}

func (this *DSU) Union(x int, y int) bool {
	px := this.Find(x)
	py := this.Find(y)

	if px == py {
		return false
	}
	if this.cnt[px] < this.cnt[py] {
		px, py = py, px
	}
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
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
