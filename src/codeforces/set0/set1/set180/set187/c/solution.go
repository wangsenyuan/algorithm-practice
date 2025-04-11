package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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
	n, m, k := readThreeNums(reader)
	volunteers := readNNums(reader, k)
	roads := make([][]int, m)
	for i := 0; i < m; i++ {
		roads[i] = readNNums(reader, 2)
	}
	s, t := readTwoNums(reader)
	return solve(n, roads, s, t, volunteers)
}

const inf = 1 << 30

func solve(n int, roads [][]int, s int, t int, volunteers []int) int {
	s--
	t--

	g := NewGraph(n, len(roads)*2)
	for _, road := range roads {
		u, v := road[0]-1, road[1]-1
		g.AddEdge(u, v)
		g.AddEdge(v, u)
	}

	pt := -1
	for i, v := range volunteers {
		if v == t+1 {
			pt = i
			break
		}
	}
	if pt < 0 {
		volunteers = append(volunteers, t+1)
	}

	set := NewDSU(n)
	que := make([]int, n)
	dist := make([]int, n)
	from := make([]int, n)
	check := func(q int) bool {
		for i := 0; i < n; i++ {
			dist[i] = inf
			from[i] = -1
		}
		set.Reset()
		var head, tail int
		for _, v := range volunteers {
			v--
			que[head] = v
			head++
			dist[v] = 0
			from[v] = v
		}
		for tail < head && set.Find(s) != set.Find(t) {
			u := que[tail]
			tail++
			for i := g.nodes[u]; i > 0; i = g.next[i] {
				v := g.to[i]
				if dist[u]+dist[v]+1 <= q {
					set.Union(from[u], from[v])
				}
				if dist[v] == inf {
					dist[v] = dist[u] + 1
					from[v] = from[u]
					que[head] = v
					head++
				}
			}
		}

		return set.Find(s) == set.Find(t)
	}

	if !check(n) {
		return -1
	}
	return sort.Search(n, check)
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

func (this *DSU) Reset() {
	for i := range this.arr {
		this.arr[i] = i
		this.cnt[i] = 1
	}
}
