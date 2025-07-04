package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	hills := make([][]int, n)
	for i := range n {
		hills[i] = readNNums(reader, 2)
	}
	m := readNum(reader)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = readNNums(reader, 2)
	}
	return solve(hills, queries)
}

type hill struct {
	id int
	x  int
	y  int
}

type point struct {
	x int
	y int
}

func (this point) cross(that point) int {
	return this.x*that.y - this.y*that.x
}

func solve(hills [][]int, queries [][]int) []int {
	n := len(hills)

	arr := make([]hill, n)
	for i, cur := range hills {
		arr[i] = hill{i + 1, cur[0], cur[1]}
	}

	slices.SortFunc(arr, func(a, b hill) int {
		return a.x - b.x
	})

	pos := make([]int, n)
	for i, cur := range arr {
		pos[i] = cur.id
	}

	stack := make([]int, n)
	var top int
	stack[top] = n - 1
	top++

	g := NewGraph(n, n)

	cross := func(i int, j int, k int) int {
		tmp := point{arr[i].x - arr[j].x, arr[i].y - arr[j].y}.cross(point{arr[j].x - arr[k].x, arr[j].y - arr[k].y})
		return -tmp
	}

	for i := n - 2; i >= 0; i-- {
		for top > 1 && cross(stack[top-2], stack[top-1], i) > 0 {
			top--
		}
		g.AddEdge(stack[top-1], i)
		stack[top] = i
		top++
	}
	h := bits.Len(uint(n))

	fa := make([][]int, n)
	dep := make([]int, n)

	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		fa[u] = make([]int, h)
		fa[u][0] = p
		for i := 1; i < h; i++ {
			fa[u][i] = fa[fa[u][i-1]][i-1]
		}
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			dep[v] = dep[u] + 1
			dfs(u, v)
		}
	}

	dfs(n-1, n-1)

	lca := func(u int, v int) int {
		if dep[u] < dep[v] {
			u, v = v, u
		}
		for i := h - 1; i >= 0; i-- {
			if dep[u]-1<<i >= dep[v] {
				u = fa[u][i]
			}
		}
		if u == v {
			return u
		}
		for i := h - 1; i >= 0; i-- {
			if fa[u][i] != fa[v][i] {
				u = fa[u][i]
				v = fa[v][i]
			}
		}
		return fa[u][0]
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		u, v := cur[0]-1, cur[1]-1
		j := lca(u, v)
		ans[i] = pos[j]
	}

	return ans
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
