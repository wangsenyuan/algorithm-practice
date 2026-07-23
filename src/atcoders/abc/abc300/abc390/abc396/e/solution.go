package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res, _, _, _ := drive(reader)
	if res == nil {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (res []int, X []int, Y []int, Z []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	X = make([]int, m)
	Y = make([]int, m)
	Z = make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &X[i], &Y[i], &Z[i])
	}
	res = solve(n, X, Y, Z)
	return
}

func solve(n int, X, Y, Z []int) []int {
	// A[X[i]] ^ A[Y[i]] = Z[i]
	// Z[i] = 确定的
	// X[i], Y[i]也是确定的
	// 考虑 x[i], y[i]是一条边, z[i]是边上的权
	// 那么如果存在一个圈, 这个圈上的权值异或和为0, 那么这个圈上的边权异或和为0
	// 如果出现了冲突 => nil
	// 假设存在的情况下, sum(A[?])最小
	// 这时候按位处理, 多的那一方, 设置为0

	adj := make([][]int, n+1)
	for i := range len(X) {
		adj[X[i]] = append(adj[X[i]], i)
		adj[Y[i]] = append(adj[Y[i]], i)
	}

	var freq [2][30]int
	var arr []int
	dp := make([]int, n+1)
	for i := range n + 1 {
		dp[i] = -1
	}
	var dfs func(u int, w int) bool
	dfs = func(u int, w int) bool {
		if dp[u] != -1 {
			return dp[u] == w
		}
		arr = append(arr, u)
		for d := range 30 {
			freq[(w>>d)&1][d]++
		}
		dp[u] = w
		for _, i := range adj[u] {
			v := X[i] ^ Y[i] ^ u
			if !dfs(v, w^Z[i]) {
				return false
			}
		}
		return true
	}
	ans := make([]int, n+1)

	for u := 1; u <= n; u++ {
		if dp[u] == -1 {

			if !dfs(u, 0) {
				return nil
			}

			for _, u := range arr {
				for d := range 30 {
					if freq[0][d] >= freq[1][d] {
						// 0的个数比1多
						if (dp[u]>>d)&1 == 1 {
							ans[u] |= 1 << d
						}
					} else {
						// 1的个数比0多
						if (dp[u]>>d)&1 == 0 {
							ans[u] |= 1 << d
						}
					}
				}
			}

			arr = arr[:0]
			for d := range 30 {
				freq[0][d] = 0
				freq[1][d] = 0
			}
		}
	}

	return ans[1:]
}
