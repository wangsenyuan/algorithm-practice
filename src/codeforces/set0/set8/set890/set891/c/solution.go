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
	res := drive(reader)
	fmt.Print(res)
}

func drive(reader *bufio.Reader) string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	roads := make([][]int, m)
	for i := range m {
		roads[i] = make([]int, 3)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1], &roads[i][2])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		var k int
		fmt.Fscan(reader, &k)
		queries[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &queries[i][j])
		}
	}
	res := solve(n, roads, queries)
	var buf bytes.Buffer
	for _, cur := range res {
		if cur {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
	}
	return buf.String()
}

type edge struct {
	id int
	u  int
	v  int
	w  int
}

type query struct {
	id int
	u  int
	v  int
	w  int
}

func solve(n int, roads [][]int, qs [][]int) []bool {

	m := len(roads)
	edges := make([]edge, m)
	for i := range m {
		edges[i] = edge{
			id: i,
			u:  roads[i][0] - 1,
			v:  roads[i][1] - 1,
			w:  roads[i][2],
		}
	}

	slices.SortFunc(edges, func(a, b edge) int {
		return a.w - b.w
	})

	var queries []query

	ans := make([]bool, len(qs))

	for i, cur := range qs {
		ans[i] = true
		for _, eid := range cur {
			eid--
			queries = append(queries, query{
				id: i,
				u:  roads[eid][0] - 1,
				v:  roads[eid][1] - 1,
				w:  roads[eid][2],
			})
		}
	}

	slices.SortFunc(queries, func(a, b query) int {
		return cmp.Or(a.w-b.w, a.id-b.id)
	})

	set := NewDSU(n)
	var pos int

	move := func(limit int) {
		for pos < m && edges[pos].w < limit {
			set.Union(edges[pos].u, edges[pos].v)
			pos++
		}
	}

	cur := queries[0].w
	pid := queries[0].id
	var cnt int
	move(cur)
	for i := 0; i < len(queries); i++ {
		if cur != queries[i].w {
			for range cnt {
				set.Rollback()
			}
			cnt = 0
			move(queries[i].w)
			cur = queries[i].w
			pid = -1
		}

		if pid != queries[i].id {
			for range cnt {
				set.Rollback()
			}
			cnt = 0
			pid = queries[i].id
		}

		if !ans[pid] {
			continue
		}

		u, v := queries[i].u, queries[i].v
		if set.Find(u) != set.Find(v) {
			set.Union(u, v)
			cnt++
		} else {
			ans[pid] = false
		}
	}

	return ans
}

type DSU struct {
	arr   []int
	cnt   []int
	stack [][2]int
}

func NewDSU(n int) *DSU {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}
	return &DSU{arr, cnt, nil}
}

func (this *DSU) Find(x int) int {
	for this.arr[x] != x {
		x = this.arr[x]
	}
	return x
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
	this.stack = append(this.stack, [2]int{px, py})
	this.cnt[px] += this.cnt[py]
	this.arr[py] = px
	return true
}

func (this *DSU) Rollback() {
	if len(this.stack) == 0 {
		return
	}
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	px, py := top[0], top[1]
	this.cnt[px] -= this.cnt[py]
	this.arr[py] = py
}
