package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	special := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &special[i])
	}
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, special, edges)
}

func solve(n int, special []int, edges [][]int) int {
	// 只能在一个special中添加
	set := NewDSU(n)

	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		set.Union(u, v)
	}

	// k := len(special)
	edgeCnt := make([]int, n)

	for _, e := range edges {
		u := e[0] - 1
		u = set.Find(u)
		edgeCnt[u]++
	}

	var res int

	var best int

	marked := make([]bool, n)

	for _, u := range special {
		u = set.Find(u - 1)
		m := set.cnt[u]
		// m * (m - 1) / 2 条边
		res += m*(m-1)/2 - edgeCnt[u]
		best = max(best, m)
		marked[u] = true
	}

	for i := range n {
		if set.Find(i) == i && !marked[i] {
			w := set.cnt[i]
			res += w * best
			res += w*(w-1)/2 - edgeCnt[i]
			best += w
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
