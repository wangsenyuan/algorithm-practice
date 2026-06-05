package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	items := make([][]int, n)
	for i := range n {
		items[i] = make([]int, 2)
		fmt.Fscan(reader, &items[i][0], &items[i][1])
	}
	return solve(m, items)
}

const inf = 1 << 60

func solve(m int, items [][]int) string {
	n := len(items)

	dp := make([][]int, n+1)

	for i := range n + 1 {
		dp[i] = make([]int, m+1)
	}

	dp[0][0] = 0
	for i, it := range items {
		copy(dp[i+1], dp[i][:it[0]])
		for p := it[0]; p <= m; p++ {
			dp[i+1][p] = max(dp[i][p], dp[i][p-it[0]]+it[1])
		}
	}

	best := dp[n][m]

	ans := make([]byte, n)
	fp := make([]int, m+1)
	fp[0] = 0
	for i := n - 1; i >= 0; i-- {
		p, v := items[i][0], items[i][1]

		var tmp1, tmp2 = -1, -1

		for x := range m + 1 {
			tmp1 = max(tmp1, dp[i][x]+fp[m-x])
			if x <= m-p {
				tmp2 = max(tmp2, dp[i][x]+v+fp[m-x-p])
			}
		}
		if tmp1 < best {
			ans[i] = 'A'
		} else if tmp2 < best {
			ans[i] = 'C'
		} else {
			ans[i] = 'B'
		}

		for j := m; j >= p; j-- {
			fp[j] = max(fp[j], fp[j-p]+v)
		}
	}

	return string(ans)
}
