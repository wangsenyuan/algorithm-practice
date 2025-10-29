package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res == 0 {
		fmt.Println("Nice work, Dima!")
	} else {
		fmt.Println(res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 4)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1], &edges[i][2], &edges[i][3])
	}
	return solve(n, edges)
}

func solve(n int, edges [][]int) int {
	set := NewDSU(n)
	check := func(l0 int, r0 int) bool {
		set.Reset()

		for _, cur := range edges {
			u, v, l, r := cur[0], cur[1], cur[2], cur[3]
			if l <= l0 && r0 <= r {
				set.Union(u-1, v-1)
			}
		}
		return set.Find(0) == set.Find(n-1)
	}

	var arr []int
	for _, cur := range edges {
		l, r := cur[2], cur[3]
		arr = append(arr, l, r)
	}
	slices.Sort(arr)
	arr = slices.Compact(arr)

	var res int

	for l, r := 0, 0; r < len(arr); r++ {
		for l <= r && !check(arr[l], arr[r]) {
			l++
		}
		if l <= r {
			res = max(res, arr[r]-arr[l]+1)
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

func (this *DSU) Reset() {
	for i := range len(this.arr) {
		this.arr[i] = i
		this.cnt[i] = 1
	}
}
