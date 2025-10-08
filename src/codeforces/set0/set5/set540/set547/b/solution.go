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
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) []int {
	n := len(a)

	type pair struct {
		first  int
		second int
	}

	arr := make([]pair, n)

	for i := range n {
		arr[i].first = a[i]
		arr[i].second = i
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(b.first-a.first, a.second-b.second)
	})

	ans := make([]int, n+1)

	set := NewDSU(n)

	for _, cur := range arr {
		i := cur.second
		if i > 0 && a[i-1] >= cur.first {
			set.Union(i-1, i)
		}
		if i+1 < n && a[i+1] >= cur.first {
			set.Union(i, i+1)
		}
		i = set.Find(i)
		ans[set.cnt[i]] = max(ans[set.cnt[i]], cur.first)
	}

	for i := n - 1; i > 0; i-- {
		ans[i] = max(ans[i], ans[i+1])
	}

	return ans[1:]
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
