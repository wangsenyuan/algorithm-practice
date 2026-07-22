package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, k, edges)
}

func solve(n, k int, edges [][]int) int {
	if k == 1 {
		return 0
	}
	set := NewDSU(n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		set.Union(u, v)
	}

	var ids []int
	for i := range n {
		if set.Find(i) == i {
			ids = append(ids, i)
		}
	}

	if len(ids) == 1 {
		return 1
	}

	ans := 1
	for i := 0; i < len(ids); i++ {
		s1 := set.cnt[ids[i]]
		ans = ans * s1 % k
	}

	// n (n - 2)
	if len(ids) > 2 {
		ans = ans * pow(n, len(ids)-2, k)
		ans %= k
	}

	return ans
}

func pow(a, b int, mod int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = res * a % mod
		}
		a = a * a % mod
		b >>= 1
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
