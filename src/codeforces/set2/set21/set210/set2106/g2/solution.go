package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	ask := func(arr []int) int {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? 1 %d", len(arr)))
		for _, x := range arr {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		fmt.Println(buf.String())
		return readNum(reader)
	}

	toggle := func(x int) {
		fmt.Printf("? 2 %d\n", x)
	}

	for range tc {
		n := readNum(reader)
		edges := make([][]int, n-1)
		for i := range n - 1 {
			edges[i] = readNNums(reader, 2)
		}

		res := solve(n, edges, ask, toggle)

		var buf bytes.Buffer
		buf.WriteString("!")

		for _, x := range res {
			buf.WriteString(fmt.Sprintf(" %d", x))
		}
		fmt.Println(buf.String())
	}
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

func solve(n int, edges [][]int, ask func(arr []int) int, toggle func(int)) []int {
	g := NewGraph(n+1, 2*n)
	for _, e := range edges {
		u, v := e[0], e[1]
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	marked := make([]bool, n+1)
	fa := make([]int, n+1)

	var dfs func(p int, u int)

	sz := make([]int, n+1)

	dfs = func(p int, u int) {
		sz[u] = 1
		fa[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p || marked[v] {
				continue
			}
			dfs(u, v)
			sz[u] += sz[v]
		}
	}

	var dfs1 func(p int, u int, w int) int

	dfs1 = func(p int, u int, w int) int {
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if v == p || marked[v] || sz[v]*2 < w {
				continue
			}
			return dfs1(u, v, w)
		}
		// u is the center
		return u
	}

	center := dfs1(0, 1, n)

	cd := make([][]int, n+1)

	var build func(prev int, node int)

	build = func(prev int, node int) {
		cd[prev] = append(cd[prev], node)
		marked[node] = true
		for i := g.nodes[node]; i > 0; i = g.next[i] {
			v := g.to[i]
			if !marked[v] {
				dfs(node, v)
				w := dfs1(node, v, sz[v])
				build(node, w)
			}
		}
	}

	build(0, center)

	check := func(arr []int, u int) int {
		l, r := 0, len(arr)
		found := false
		for l < r {
			mid := (l + r) / 2
			s1 := ask(arr[l : mid+1])
			toggle(u)
			s2 := ask(arr[l : mid+1])
			if abs(s2-s1) == 2*(mid-l+1) {
				// 肯定在外部
				l = mid + 1
			} else {
				found = true
				r = mid
			}
		}

		if found {
			return arr[l]
		}

		return -1
	}

	var findRoot func(p int, u int) int

	findRoot = func(p int, u int) int {
		if len(cd[u]) == 0 {
			return u
		}

		var tmp []int

		for _, v := range cd[u] {
			if v != p {
				tmp = append(tmp, v)
			}
		}

		v := check(tmp, u)
		if v < 0 {
			return u
		}
		return findRoot(u, v)
	}

	root := findRoot(0, center)

	ans := make([]int, n+1)

	var dfs4 func(p int, u int, sum int)

	dfs4 = func(p int, u int, sum int) {
		ans[u] = ask([]int{u}) - sum
		sum += ans[u]
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dfs4(u, v, sum)
			}
		}
	}

	dfs4(0, root, 0)

	return ans[1:]
}

func abs(num int) int {
	return max(num, -num)
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
