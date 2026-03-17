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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	edges := make([][]int, m)
	for i := range m {
		edges[i] = make([]int, 2)
		fmt.Fscan(reader, &edges[i][0], &edges[i][1])
	}
	return solve(n, edges)
}

func solve(n int, roads [][]int) int {
	set := NewDSU(n)
	for _, cur := range roads {
		u, v := cur[0]-1, cur[1]-1
		set.Union(u, v)
	}
	freq := make([]int, n+1)
	for i := range n {
		if set.Find(i) == i {
			freq[set.cnt[i]]++
		}
	}
	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = inf
	}

	dp[0] = 0

	play := func(s int, a int, c int) {
		k := 1
		for 1<<k <= c+1 {
			k++
		}
		// 1 << k - 1 <= c
		k--
		var arr []int
		for i := range k {
			arr = append(arr, 1<<i)
		}
		if c > (1<<k)-1 {
			arr = append(arr, c-(1<<k)+1)
		}
		for _, v := range arr {
			for i := s; i >= v*a; i-- {
				dp[i] = min(dp[i], dp[i-v*a]+v)
			}
		}
	}

	var sum int

	for i := 1; i <= n; i++ {
		sum += i * freq[i]
		if freq[i] > 0 {
			play(sum, i, freq[i])
		}
	}

	ans := inf

	for i := 1; i <= n; i++ {
		if dp[i] < inf && checkLucky(i) {
			ans = min(ans, dp[i])
		}
	}
	if ans == inf {
		return -1
	}

	return ans - 1
}

func checkLucky(x int) bool {
	for x > 0 {
		r := x % 10
		if r != 4 && r != 7 {
			return false
		}
		x /= 10
	}
	return true
}

const inf = 1 << 60

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
