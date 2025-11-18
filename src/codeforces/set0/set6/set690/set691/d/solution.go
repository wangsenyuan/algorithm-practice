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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1:len(s) - 1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	pairs := make([][]int, m)
	for i := range m {
		var u, v int
		fmt.Fscan(reader, &u, &v)
		pairs[i] = []int{u, v}
	}
	return solve(p, m, pairs)
}

func solve(p []int, m int, pairs [][]int) []int {
	n := len(p)

	set := NewDSU(n)
	for _, cur := range pairs {
		u, v := cur[0]-1, cur[1]-1
		set.Union(u, v)
	}

	pos := make([][]int, n)
	nums := make([][]int, n)
	for i := range n {
		j := set.Find(i)
		pos[j] = append(pos[j], i)
		nums[j] = append(nums[j], p[i])
	}
	res := make([]int, n)

	for i := range n {
		slices.Sort(nums[i])
		slices.Reverse(nums[i])
		for u, j := range pos[i] {
			res[j] = nums[i][u]
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
