package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	x, res := process(reader)
	fmt.Println(x, len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (int, []int) {
	n := readNum(reader)
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = readNNums(reader, 3)
	}
	return solve(n, edges)
}

type edge struct {
	u  int
	v  int
	w  int
	id int
}

func solve(n int, edges [][]int) (int, []int) {
	es := make([]edge, n-1)
	g := NewGraph(n, 2*n)
	for i, cur := range edges {
		u, v, w := cur[0]-1, cur[1]-1, cur[2]
		es[i] = edge{u, v, w, i}
		g.AddEdge(u, v, i)
		g.AddEdge(v, u, i)
	}

	fa := make([]int, n)
	dep := make([]int, n-1)

	var dfs func(p int, u int, d int)
	dfs = func(p int, u int, d int) {
		fa[u] = p
		for i := g.nodes[u]; i > 0; i = g.next[i] {
			v := g.to[i]
			if p != v {
				dep[g.val[i]] = d
				dfs(u, v, d+1)
			}
		}
	}

	dfs(0, 0, 0)

	// 在相同长度时，先处理深的那些边
	slices.SortFunc(es, func(a, b edge) int {
		return cmp.Or(a.w-b.w, dep[b.id]-dep[a.id])
	})

	ans := make([]int, n-1)

	set := NewDSU(n)
	cnt := make([]int, n-1)

	for i := 0; i < n-1; {
		j := i
		for i < n-1 && es[i].w == es[j].w {
			u, v := es[i].u, es[i].v
			if fa[u] == v {
				cnt[i] = set.cnt[set.Find(u)]
			} else {
				cnt[i] = set.cnt[set.Find(v)]
			}
			set.Union(es[i].u, es[i].v)
			i++
		}

		for j < i {
			u := es[j].u
			tmp := set.cnt[set.Find(u)]
			ans[j] = cnt[j] * (tmp - cnt[j])
			j++
		}
	}

	x := slices.Max(ans)

	var res []int
	for i, v := range ans {
		if x == v {
			res = append(res, es[i].id+1)
		}
	}
	sort.Ints(res)
	return x * 2, res
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
