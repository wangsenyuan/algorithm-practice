package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	total, ans := drive(reader)

	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", total))
	for _, v := range ans {
		buf.WriteString(fmt.Sprintf("%d %d\n", v[0], v[1]))
	}
	buf.WriteTo(os.Stdout)
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

func drive(reader *bufio.Reader) (total int, ans [][]int) {
	first := readNNums(reader, 4)
	n, m, k, w := first[0], first[1], first[2], first[3]
	a := make([][]string, k)

	for i := range k {
		a[i] = make([]string, n)
		for j := range n {
			a[i][j] = readString(reader)
		}
	}
	return solve(n, m, k, w, a)
}

func solve(n int, m int, k int, w int, levels [][]string) (total int, ans [][]int) {
	calcDiff := func(l int, r int) int {
		var sum int
		for i := range n {
			for j := range m {
				if levels[l][i][j] != levels[r][i][j] {
					sum++
				}
			}
		}
		return sum
	}

	type edge struct {
		u int
		v int
		w int
	}

	var edges []edge

	for i := range k {
		for j := range i {
			diff := calcDiff(j, i)
			if diff*w < n*m {
				edges = append(edges, edge{j, i, diff * w})
			}
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return a.w - b.w
	})

	set := NewDSU(k)

	g := NewGraph(k, 2*len(edges))

	for _, e := range edges {
		if set.Union(e.u, e.v) {
			g.AddEdge(e.u, e.v)
			g.AddEdge(e.v, e.u)
			total += e.w
		}
	}

	marked := make([]bool, k)
	que := make([]int, k)

	bfs := func(s int) {
		var head, tail int
		que[head] = s
		head++
		marked[s] = true
		total += n * m
		ans = append(ans, []int{s + 1, 0})
		for tail < head {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if !marked[v] {
					marked[v] = true
					ans = append(ans, []int{v + 1, u + 1})
					que[head] = v
					head++
				}
			}
		}
	}

	for u := range k {
		if !marked[u] {
			bfs(u)
		}
	}
	return
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
