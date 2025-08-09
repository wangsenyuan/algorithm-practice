package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, d int
	fmt.Fscan(reader, &n, &d)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(d, p)
}

func solve(d int, p []int) int {
	n := p[len(p)-1] + 1
	gems := make([]int, n)
	for _, i := range p {
		gems[i]++
	}
	dp := make([][]int, n)
	offset := 250

	for i := range n {
		dp[i] = make([]int, 2*offset)
		for j := range 2 * offset {
			dp[i][j] = -inf
		}
	}

	var dfs func(i int, j int) int
	dfs = func(i int, j int) int {
		if i >= n {
			return 0
		}
		j1 := j - (d - offset)
		if dp[i][j1] >= 0 {
			return dp[i][j1]
		}
		var res int
		if j == 1 {
			res = gems[i] + max(dfs(i+j, j), dfs(i+j+1, j+1))
		} else {
			res = gems[i] + max(dfs(i+j, j), dfs(i+j-1, j-1), dfs(i+j+1, j+1))
		}

		dp[i][j1] = res
		return res
	}

	return dfs(d, d)
}

const inf = 1 << 60
