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
	var n, m, w int
	fmt.Fscan(reader, &n, &m, &w)
	weights := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &weights[i])
	}
	beauties := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &beauties[i])
	}
	friends := make([][]int, m)
	for i := range m {
		friends[i] = make([]int, 2)
		fmt.Fscan(reader, &friends[i][0], &friends[i][1])
	}
	return solve(w, weights, beauties, friends)
}

const inf = 1 << 60

func solve(w int, weights []int, beauties []int, friends [][]int) int {
	n := len(weights)

	set := NewDSU(n)

	for _, cur := range friends {
		u, v := cur[0]-1, cur[1]-1
		set.Union(u, v)
	}

	groups := make([][]int, n)

	for i := range n {
		j := set.Find(i)
		groups[j] = append(groups[j], i)
	}

	// max beauties for w
	dp := make([]int, w+1)
	ndp := make([]int, w+1)
	for i := range w + 1 {
		dp[i] = -inf
		ndp[i] = -inf
	}
	dp[0] = 0

	for _, cur := range groups {
		if len(cur) == 0 {
			continue
		}
		sum := make([]int, 2)
		for _, i := range cur {
			x, b := weights[i], beauties[i]
			sum[0] += x
			sum[1] += b
			for j := w; j >= x; j-- {
				ndp[j] = max(ndp[j], dp[j-x]+b)
			}
		}
		// 或者全部邀请
		for j := w; j >= sum[0]; j-- {
			ndp[j] = max(ndp[j], dp[j-sum[0]]+sum[1])
		}
		for j := range w + 1 {
			dp[j] = max(dp[j], ndp[j])
			ndp[j] = -inf
		}
	}

	return slices.Max(dp)
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
