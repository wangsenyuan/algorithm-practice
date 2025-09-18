package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	var buf bytes.Buffer
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d %d\n", cur[0], cur[1], cur[2]))
	}

	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) (nodes [][]int, edges [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	nodes = make([][]int, n)
	for i := range n {
		var c, s int
		fmt.Fscan(reader, &c, &s)
		nodes[i] = []int{c, s}
	}
	edges = solve(n, nodes)
	return
}

type pair struct {
	first  int
	second int
}

func solve(n int, nodes [][]int) [][]int {
	var red []pair
	var blue []pair
	for i := range n {
		if nodes[i][0] == 0 {
			red = append(red, pair{nodes[i][1], i})
		} else {
			blue = append(blue, pair{nodes[i][1], i})
		}
	}

	cmp := func(a, b pair) int {
		return cmp.Or(a.first-b.first, a.second-b.second)
	}

	slices.SortFunc(red, cmp)
	slices.SortFunc(blue, cmp)
	var res [][]int
	var i, j int

	set := NewDSU(n)

	for i < len(red) && j < len(blue) {
		w := min(red[i].first, blue[j].first)
		set.Union(red[i].second, blue[j].second)
		res = append(res, []int{red[i].second + 1, blue[j].second + 1, w})
		red[i].first -= w
		blue[j].first -= w
		if red[i].first == 0 {
			i++
		}
		if blue[j].first == 0 {
			j++
		}
	}
	// 没有连起来， 使用0去连接
	u := red[0].second
	for _, cur := range blue {
		if set.Union(u, cur.second) {
			res = append(res, []int{u + 1, cur.second + 1, 0})
		}
	}
	v := blue[0].second
	for _, cur := range red {
		if set.Union(v, cur.second) {
			res = append(res, []int{cur.second + 1, v + 1, 0})
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
