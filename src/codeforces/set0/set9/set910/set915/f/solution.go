package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	edges := make([][]int, n-1)
	for i := range n - 1 {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(a, edges)
}

type node struct {
	id  int
	val int
}

func solve(a []int, edges [][]int) int64 {
	n := len(a)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	arr := make([]node, n)
	for i, v := range a {
		arr[i] = node{i, v}
	}

	slices.SortFunc(arr, func(a node, b node) int {
		return cmp.Or(a.val-b.val, a.id-b.id)
	})

	set := NewDSU(n)

	var res int
	for _, cur := range arr {
		u := cur.id
		sz := 1

		for _, v := range adj[u] {
			if (a[v] < a[u] || a[v] == a[u] && v < u) && set.Find(v) != set.Find(u) {
				s1 := set.cnt[set.Find(v)]
				res += cur.val * s1 * sz
				sz += s1
				set.Union(v, u)
			}
		}
	}

	slices.SortFunc(arr, func(a node, b node) int {
		return cmp.Or(b.val-a.val, a.id-b.id)
	})

	set.Reset()

	for _, cur := range arr {
		u := cur.id
		sz := 1
		for _, v := range adj[u] {
			if (a[v] > a[u] || a[v] == a[u] && v < u) && set.Find(v) != set.Find(u) {
				s1 := set.cnt[set.Find(v)]
				res -= cur.val * s1 * sz
				sz += s1
				set.Union(v, u)
			}
		}
	}

	return int64(res)
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
