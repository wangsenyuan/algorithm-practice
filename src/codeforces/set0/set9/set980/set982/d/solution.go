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
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) int {
	// k = 最小值是一个答案吗？
	// 不行， 因为有non_empty的要求
	n := len(a)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{a[i], i}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	set := NewDSU(n)

	var mx_sz int
	// 要最大化locations
	var locations int
	var cnt int

	res := arr[n-1].first + 1

	for i, cur := range arr {
		j := cur.second
		// j还没有被处理
		if j > 0 && a[j-1] < cur.first {
			set.Union(j-1, j)
			cnt--
		}
		if j+1 < n && a[j+1] < cur.first {
			set.Union(j, j+1)
			cnt--
		}
		mx_sz = max(mx_sz, set.cnt[set.Find(j)])
		cnt++

		if cnt > locations && cnt*mx_sz == i+1 {
			locations = cnt
			res = cur.first + 1
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
