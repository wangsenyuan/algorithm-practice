package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for i, x := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, x)
	}
	fmt.Fprintln(writer)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	groups := make([][]int, m)
	for i := range m {
		var k int
		fmt.Fscan(reader, &k)
		groups[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &groups[i][j])
		}
	}
	return solve(n, groups)
}

func solve(n int, groups [][]int) []int {
	set := NewDSU(n)
	for _, cur := range groups {
		for i := 1; i < len(cur); i++ {
			set.Union(cur[i-1]-1, cur[i]-1)
		}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		w := set.Find(i)
		ans[i] = set.cnt[w]
	}
	return ans
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
