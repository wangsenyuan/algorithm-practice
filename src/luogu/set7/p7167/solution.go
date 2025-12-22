package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)

	plates := make([][]int, n)
	for i := 0; i < n; i++ {
		plates[i] = readNNums(reader, 2)
	}

	solver := NewSolver(plates)

	var buf bytes.Buffer
	for i := 0; i < m; i++ {
		p, x := readTwoNums(reader)
		res := solver.Query(p, x)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

type Solver struct {
	n    int
	deps []int
	P    [][]int
	vols []int
}

const H = 18

func NewSolver(cups [][]int) *Solver {
	n := len(cups)
	next := make([]int, n+1)
	stack := make([]int, n)
	var top int
	for i := n - 1; i >= 0; i-- {
		next[i] = n
		for top > 0 && cups[stack[top-1]][0] <= cups[i][0] {
			top--
		}
		if top > 0 {
			next[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}

	g := NewGraph(n+1, n)

	for i := 0; i < n; i++ {
		g.AddEdge(next[i], i)
	}

	vol := make([]int, n+1)

	P := make([][]int, n+1)

	D := make([]int, n+1)

	var dfs func(p int, u int)

	dfs = func(p int, u int) {
		if u < n {
			vol[u] = vol[p] + cups[u][1]
		}

		P[u] = make([]int, H)
		P[u][0] = p
		for i := 1; i < H; i++ {
			P[u][i] = P[P[u][i-1]][i-1]
		}

		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			D[v] = D[u] + 1
			dfs(u, v)
		}
	}
	P[n] = make([]int, H)

	dfs(n, n)

	return &Solver{n, D, P, vol}
}

func (s *Solver) Query(p int, v int) int {
	p--
	if s.vols[p] < v {
		return 0
	}
	// sum >= v
	// 在路径 p -> next[p] -> next[next[p]] ... -> n
	// 存在某个i, vols[p] - vols[i] >= v
	// 那么那个i就是答案
	// vols[i] <= vols[p] - v
	getKthAnc := func(u int, k int) int {
		for i := H - 1; i >= 0; i-- {
			if (k>>i)&1 == 1 {
				u = s.P[u][i]
				k -= 1 << i
			}
		}
		return u
	}

	check := func(d int) bool {
		i := getKthAnc(p, s.deps[p]-d)
		// d 越大， 离p越近，它们的差值越小
		return s.vols[p]-s.vols[i] < v
	}

	d := search(s.deps[p], check)

	return getKthAnc(p, s.deps[p]-d) + 1
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

type Graph struct {
	nodes []int
	next  []int
	to    []int
	cur   int
}

func NewGraph(n int, e int) *Graph {
	nodes := make([]int, n)
	e++
	next := make([]int, e)
	to := make([]int, e)
	return &Graph{nodes, next, to, 0}
}

func (g *Graph) AddEdge(u, v int) {
	g.cur++
	g.next[g.cur] = g.nodes[u]
	g.nodes[u] = g.cur
	g.to[g.cur] = v
}
