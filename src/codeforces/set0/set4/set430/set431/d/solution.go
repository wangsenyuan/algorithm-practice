package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, k int
	fmt.Scanf("%d %d", &m, &k)
	n := solve(m, k)
	fmt.Println(n)
}

const inf = 1e18

const H = 64

var C [H][H]int

func init() {
	C[0][0] = 1
	for i := 1; i < H; i++ {
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}
}

func solve(m int, k int) int {
	if m == 0 {
		if k > 1 {
			return 1
		}
		// k == 1
		return 1 << 60
	}
	check := func(n int) bool {
		if n == 0 {
			return false
		}
		return count(n*2, k)-count(n, k) >= m
	}
	return sort.Search(inf, check)
}

func count(n int, k int) int {
	var cnt int
	var ans int
	for i := H - 1; i >= 0; i-- {
		if n&(1<<i) > 0 {
			ans += C[i][k-cnt]
			cnt++
		}
		if cnt == k {
			ans++
			break
		}
	}

	return ans
}
