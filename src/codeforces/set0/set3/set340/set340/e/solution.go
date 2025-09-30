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

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	p := readNNums(reader, n)
	return solve(p)
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(nums ...int) int {
	res := 1
	for _, v := range nums {
		res *= v
		res %= mod
	}
	return res
}

func solve(p []int) int {
	n := len(p)
	var free int
	marked := make([]bool, n)
	var m int
	for i := 0; i < n; i++ {
		if i+1 == p[i] {
			return 0
		}
		if p[i] < 0 {
			m++
		}
		if p[i] == -1 || marked[i] {
			continue
		}
		j := i
		for !marked[j] {
			marked[j] = true
			j = p[j]
			if j < 0 {
				// 这里遇到了一个free的数字
				free++
				break
			}
			j--
		}
	}

	F := make([]int, n+1)
	F[0] = 1
	for i := 1; i <= n; i++ {
		F[i] = mul(i, F[i-1])
	}

	// dp[i][j] 是共有i个位置，其中有j个free的数时的计数
	// dp[i][0] = f[i]
	// dp[i][i] = 1
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = -1
		}
	}

	var dfs func(u int, v int) int
	dfs = func(u int, v int) (res int) {
		if u < v {
			return 0
		}

		if u == v {
			return F[u]
		}

		if dp[u][v] >= 0 {
			return dp[u][v]
		}
		defer func() {
			dp[u][v] = res
		}()

		if v == 0 {
			// 最后一个数字放在除最后一个位置上, 这时会多出一个free的数字
			res = mul(u-1, dfs(u-1, v+1))
		} else {
			// 将一个free的数字放在free数字对应的位置上，或者是其他的位置上
			res = add(mul(v, dfs(u-1, v-1)), mul(u-v, dfs(u-1, v)))
		}

		return
	}

	return dfs(m, free)
}
