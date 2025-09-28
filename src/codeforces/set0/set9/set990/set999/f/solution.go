package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	c := readNNums(reader, n*k)
	f := readNNums(reader, n)
	h := readNNums(reader, k)
	return solve(n, k, c, f, h)
}

const inf = 1 << 60

func solve(n int, k int, c []int, f []int, h []int) int {
	// 假设两个人， i, j, 他们都喜欢x， 且此时分配给 a[i], a[j]是，i, j拿到x的数量，且a[i] > a[j]
	// 也就是说，目前的分数是 h[a[i]] + h[a[j]]
	// 现在有一张新的x，如果分配给i, h[a[i] + 1] - h[a[i]], 如果分配给j, h[a[j] + 1] - h[a[j]]
	// 两个的差值 = h[a[i] + 1] - h[a[i]] - (h[a[j]] + 1) - h[a[j]])
	// = h[a[i] + 1] - h[a[j] + 1] - (h[a[i]] - h[a[j]])
	// 所以不大对
	// 貌似要计算把m个x，分配给u个人， 且每个人分配不超过k,时的最优解
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, i*k+1)
		for j := range i*k + 1 {
			dp[i][j] = -inf
		}
	}

	dp[0][0] = 0

	for i := range n {
		for j := range i*k + 1 {
			dp[i+1][j] = max(dp[i+1][j], dp[i][j])
			for x := 1; x <= k; x++ {
				dp[i+1][j+x] = max(dp[i+1][j+x], dp[i][j]+h[x-1])
			}
		}
	}

	sort.Ints(c)
	sort.Ints(f)

	var res int

	for i, j := 0, 0; i < len(c); {
		i1 := i
		for i < len(c) && c[i] == c[i1] {
			i++
		}
		for j < len(f) && f[j] < c[i1] {
			j++
		}
		if j < len(f) && f[j] == c[i1] {
			j1 := j
			for j < len(f) && f[j] == c[i1] {
				j++
			}
			cnt2 := j - j1
			cnt1 := min(cnt2*k, i-i1)
			res += dp[cnt2][cnt1]
		}
	}

	return res
}
