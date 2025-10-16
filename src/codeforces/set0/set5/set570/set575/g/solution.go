package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	min_time, path := drive(reader)
	fmt.Fprintln(writer, min_time)
	fmt.Fprintln(writer, len(path))
	for _, v := range path {
		fmt.Fprint(writer, v, " ")
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) (string, []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 3)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2])
	}
	return solve(n, edges)
}

type pair struct {
	first  int
	second int
}

func solve(n int, edges [][]int) (min_time string, path []int) {
	m := len(edges)
	g := NewGraph(n, 2*m)
	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		g.AddEdge(u, v, w)
		g.AddEdge(v, u, w)
	}
	next := make([]int, n)
	d0 := make([]int, n)
	d1 := make([]int, n)
	for i := range n {
		d0[i] = -1
		d1[i] = -1
		next[i] = -1
	}

	que := make([]int, n)
	var head, tail int

	d1[n-1] = 0
	que[head] = n - 1
	head++
	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			w := g.val[i]
			if w == 0 && d1[v] < 0 {
				next[v] = u
				d1[v] = d1[u] + 1
				que[head] = v
				head++
			}
		}
	}

	head = 0
	tail = 0
	d0[0] = 0
	que[head] = 0
	head++
	for tail < head {
		u := que[tail]
		tail++
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if d0[v] < 0 {
				d0[v] = d0[u] + 1
				que[head] = v
				head++
			}
		}
	}

	mn := n
	for i := range n {
		if d1[i] >= 0 {
			mn = min(mn, d0[i])
		}
	}

	var arr []pair
	for i := range n {
		if d1[i] >= 0 && d0[i] == mn {
			arr = append(arr, pair{d1[i], i})
		}
	}

	// 这个sort有点不理解
	slices.SortFunc(arr, func(x pair, y pair) int {
		return x.first - y.first
	})

	check := func(arr []pair) bool {
		// 是否已经访问到了0号节点
		for _, v := range arr {
			if v.second == 0 {
				return true
			}
		}
		return false
	}

	for !check(arr) {
		mn := 10
		for _, cur := range arr {
			u := cur.second
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if d0[v]+1 == d0[u] {
					mn = min(mn, g.val[i])
				}
			}
		}

		var tmp []pair
		for _, cur := range arr {
			u := cur.second
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if mn == g.val[i] && d0[v]+1 == d0[u] && next[v] < 0 {
					next[v] = u
					tmp = append(tmp, pair{d0[v], v})
				}
			}
		}
		arr = tmp
	}

	var time []int

	var u int
	for u != n-1 {
		path = append(path, u)
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == next[u] {
				time = append(time, g.val[i])
				break
			}
		}
		u = next[u]
	}
	path = append(path, n-1)

	slices.Reverse(time)

	for len(time) > 0 && time[0] == 0 {
		time = time[1:]
	}
	if len(time) == 0 {
		time = append(time, 0)
	}

	var buf strings.Builder
	for _, v := range time {
		buf.WriteString(strconv.Itoa(v))
	}

	min_time = buf.String()

	return
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
