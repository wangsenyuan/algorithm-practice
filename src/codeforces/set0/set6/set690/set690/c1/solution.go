package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	connectors := make([][]int, m)
	for i := range m {
		connectors[i] = make([]int, 2)
		fmt.Fscan(reader, &connectors[i][0], &connectors[i][1])
	}
	return solve(n, connectors)
}

func solve(n int, connectors [][]int) bool {
	m := len(connectors)
	if m != n-1 {
		return false
	}
	set := NewDSU(n)
	for _, cur := range connectors {
		u, v := cur[0]-1, cur[1]-1
		set.Union(u, v)
	}

	root := set.Find(0)
	return set.cnt[root] == n
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
