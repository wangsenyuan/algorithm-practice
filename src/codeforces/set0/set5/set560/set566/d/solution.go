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
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	q := make([][]int, m)
	for i := range m {
		q[i] = make([]int, 3)
		fmt.Fscan(reader, &q[i][0], &q[i][1], &q[i][2])
	}
	return solve(n, q)
}

func solve(n int, q [][]int) []string {
	set := NewDSU(n)

	f := make([]int, n)
	for i := range n {
		f[i] = i + 1
	}

	update := func(l int, r int) {
		for l < r {
			i := f[l]
			if i > r {
				break
			}
			set.Union(l, i)
			f[l] = max(f[l], f[r])
			l = i
		}
	}

	var ans []string
	for _, cur := range q {
		u, v := cur[1]-1, cur[2]-1
		if u > v {
			u, v = v, u
		}
		switch cur[0] {
		case 1:
			if set.Union(u, v) {
				if u+1 == v {
					f[u] = f[v]
				}
			}
		case 2:
			update(u, v)
		default:
			if set.Find(u) == set.Find(v) {
				ans = append(ans, "YES")
			} else {
				ans = append(ans, "NO")
			}
		}
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
