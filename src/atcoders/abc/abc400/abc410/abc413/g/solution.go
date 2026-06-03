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
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func drive(reader *bufio.Reader) bool {
	var h, w, n int
	fmt.Fscan(reader, &h, &w, &n)
	obstacles := make([][]int, n)
	for i := range n {
		obstacles[i] = make([]int, 2)
		fmt.Fscan(reader, &obstacles[i][0], &obstacles[i][1])
	}
	return solve(h, w, obstacles)
}

type pair struct {
	first  int
	second int
}

func solve(h int, w int, obstacles [][]int) bool {
	n := len(obstacles)

	id := make(map[pair]int)
	for i, cur := range obstacles {
		id[pair{cur[0], cur[1]}] = i
	}
	// n, n+1, n+2, n+3, 上下左右
	set := NewDSU(n + 4)
	for i, cur := range obstacles {
		r, c := cur[0], cur[1]
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				nr, nc := r+dx, c+dy
				if nr == 0 && (nc == 0 || nc == w+1) || nr == h+1 && (nc == 0 || nc == w+1) {
					// 角落不处理
					continue
				}
				if nr == 0 {
					set.Union(i, n)
				} else if nr == h+1 {
					set.Union(i, n+2)
				} else if nc == 0 {
					set.Union(i, n+3)
				} else if nc == w+1 {
					set.Union(i, n+1)
				} else if i1, ok := id[pair{nr, nc}]; ok {
					set.Union(i, i1)
				}
			}
		}
	}

	blocks := [][]int{
		{n, n + 2},     // 上下
		{n + 1, n + 3}, // 左右
		{n, n + 3},     // 上左
		{n + 1, n + 2}, // 下右
	}

	for _, cur := range blocks {
		u, v := cur[0], cur[1]
		if set.Find(u) == set.Find(v) {
			return false
		}
	}

	return true
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
