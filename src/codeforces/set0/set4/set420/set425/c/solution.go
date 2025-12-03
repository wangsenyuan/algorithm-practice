package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) int {
	var n, m, s, e int
	fmt.Fscan(reader, &n, &m, &s, &e)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	return solve(s, e, a, b)
}

const inf = 1 << 60

func solve(s int, e int, a []int, b []int) int {

	// 这个e很大，1e3, 那么 s/e <= 100
	// 也就是说最多选择100个数
	if s < e {
		return 0
	}
	n := len(a)
	m := len(b)
	k := s / e

	const X = 100000

	pos := make([][]int, X+1)
	for i := range m {
		pos[b[i]] = append(pos[b[i]], i+1)
	}

	dp := make([]int, k+1)
	ndp := make([]int, k+1)

	for i := range k + 1 {
		dp[i] = inf
		ndp[i] = inf
	}
	dp[0] = 0

	var best int

	for i := range n {
		for x := range k {
			w := dp[x]
			// 在w的后面找到最近的a[i]`
			j1 := sort.SearchInts(pos[a[i]], w+1)
			if j1 < len(pos[a[i]]) {
				ndp[x+1] = min(ndp[x+1], pos[a[i]][j1])
			}
		}
		for x := range k + 1 {
			dp[x] = min(dp[x], ndp[x])
			if dp[x] <= m && x*e+i+1+dp[x] <= s {
				best = max(best, x)
			}
			ndp[x] = inf
		}
	}

	return best
}
