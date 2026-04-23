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
	var m, n int
	fmt.Fscan(reader, &m, &n)
	a := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	A := make([][]int, m)
	for i := range m {
		var k int
		fmt.Fscan(reader, &k)
		A[i] = make([]int, k)
		for j := range k {
			fmt.Fscan(reader, &A[i][j])
		}
	}
	return solve(A, a, b)
}

type data struct {
	id  int
	val int
}

func solve(A [][]int, a []int, b []int) int {
	m := len(a)
	n := len(b)
	// 构造n+m个点
	var tot int
	var arr []data
	for i, cur := range A {
		for _, j := range cur {
			tot += a[i] + b[j-1]
			arr = append(arr, data{id: i*n + j - 1, val: a[i] + b[j-1]})
		}
	}

	slices.SortFunc(arr, func(a, b data) int {
		return b.val - a.val
	})

	set := NewDSU(n + m)
	for _, cur := range arr {
		i, j := cur.id/n, cur.id%n
		i += n
		if set.Union(i, j) {
			tot -= cur.val
		}
	}
	return tot
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
