package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := drive(reader)
	fmt.Println(ans)
}

func drive(reader *bufio.Reader) int {
	var a, n, m int
	fmt.Fscan(reader, &a, &n, &m)
	rain_segments := make([][]int, n)
	for i := range n {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		rain_segments[i] = []int{l, r}
	}
	umbrellas := make([][]int, m)
	for i := range m {
		var x, p int
		fmt.Fscan(reader, &x, &p)
		umbrellas[i] = []int{x, p}
	}
	return solve(a, rain_segments, umbrellas)
}

const inf = 1 << 60

func solve(a int, rain_segments [][]int, umbrellas [][]int) int {

	flag := make([]int, a+1)
	for _, cur := range rain_segments {
		l := cur[0]
		r := cur[1]
		flag[l]++
		flag[r]--
	}
	for i := 1; i <= a; i++ {
		flag[i] += flag[i-1]
	}

	m := len(umbrellas)
	dp := make([]int, m+1)
	ndp := make([]int, m+1)
	for i := range m + 1 {
		dp[i] = inf
		ndp[i] = inf
	}

	dp[0] = 0

	for i := range a {
		best := slices.Min(dp)
		for j := range m {
			x, p := umbrellas[j][0], umbrellas[j][1]
			if x < i {
				// 仍然拿在手里
				ndp[j+1] = min(ndp[j+1], dp[j+1]+p)
			} else if x == i {
				// 第一次拿到手里
				ndp[j+1] = min(ndp[j+1], best+p)
			}
		}
		if flag[i] == 0 {
			ndp[0] = best
		} else {
			ndp[0] = inf
		}
		copy(dp, ndp)
		for j := range m + 1 {
			ndp[j] = inf
		}
	}

	res := slices.Min(dp)
	if res == inf {
		return -1
	}
	return res
}
