package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (segments [][]int, res string) {
	var n int
	fmt.Fscan(reader, &n)
	segments = make([][]int, n)
	for i := range n {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		segments[i] = []int{l, r}
	}
	res = solve(n, segments)
	return
}

func solve(n int, segments [][]int) string {
	dp := make([][]bool, n)
	for i := range n {
		dp[i] = make([]bool, n)
	}

	for j := range n {
		for i := j; i >= 0; i-- {
			l, r := segments[i][0], segments[i][1]
			if i == j {
				if l <= 1 {
					dp[i][j] = true
				}
			} else {
				// i < j
				// 这个是准长度
				w := (j - i + 1) * 2
				if w-1 >= l && w-1 <= r && dp[i+1][j] {
					dp[i][j] = true
					continue
				}

				for k := i; k < j; k++ {
					if dp[i][k] && dp[k+1][j] {
						dp[i][j] = true
						break
					}
				}
			}
		}
	}
	if !dp[0][n-1] {
		return "IMPOSSIBLE"
	}

	buf := make([]byte, 2*n)

	var dfs func(pos int, l int, r int)
	dfs = func(pos int, l int, r int) {
		w := (r - l + 1) * 2

		// dp[l][r] must be treu
		if w-1 >= segments[l][0] && w-1 <= segments[l][1] && (w == 2 || dp[l+1][r]) {
			buf[pos] = '('
			buf[pos+w-1] = ')'
			if l < r {
				dfs(pos+1, l+1, r)
			}
			return
		}
		for k := l; k < r; k++ {
			if dp[l][k] && dp[k+1][r] {
				dfs(pos, l, k)
				dfs(pos+(k-l+1)*2, k+1, r)
				break
			}
		}
	}

	dfs(0, 0, n-1)

	return string(buf)
}
