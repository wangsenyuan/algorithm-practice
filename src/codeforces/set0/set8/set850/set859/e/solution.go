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
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, 2)
		fmt.Fscan(reader, &a[i][0], &a[i][1])
	}
	return solve(n, a)
}

const mod = 1000000007

func mul(a, b int) int {
	return a * b % mod
}

func add(a, b int) int {
	return (a + b) % mod
}

func solve(n int, a [][]int) int {
	n2 := 2 * n
	set := NewDSU(n2)

	for _, cur := range a {
		u, v := cur[0]-1, cur[1]-1
		set.Union(u, v)
	}

	edgeCnt := make([]int, n2)

	special := make([]bool, n2)

	for _, cur := range a {
		u := cur[0] - 1
		v := cur[1] - 1
		p := set.Find(u)
		if u == v {
			special[p] = true
		}
		edgeCnt[p]++
	}

	res := 1

	for i := range n2 {
		j := set.Find(i)
		if i == j {
			e := edgeCnt[i]
			v := set.cnt[i]
			if e+1 == v {
				res = mul(res, v)
			} else if !special[i] {
				res = mul(res, 2)
			}
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
