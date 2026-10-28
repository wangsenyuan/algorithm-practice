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
	var n, k int
	var s string
	fmt.Fscan(reader, &n, &k, &s)
	return solve(n, k, s)
}

const mod = 1000000007

func add(a int, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a int, b int) int {
	return int(int64(a) * int64(b) % mod)
}

func solve(n int, k int, s string) int {
	// dp[i][j]：填好前 i 位，且第 i 位与 s 不同，已经确定的 beauty 是 j。
	// 上一次不同在 p 时，(p, i) 必须与 s 相同。
	// 这一位更小则 beauty 不变；更大则增加 (i-p)*(n-i+1)。
	// 最后一次不同之后的后缀全部与 s 相同，不再增加 beauty。
	dp := make([][]int, n+1)
	// pref := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, k+1)
		// pref[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	pref := make([]int, k+1)
	pref[0] = 1
	for i := 1; i <= n; i++ {
		w := int(s[i-1] - 'a')
		endings := n - i + 1
		for j := range k + 1 {
			dp[i][j] = mul(w, pref[j])
			if w == 25 || endings > j {
				continue
			}
			maxLen := min(j/endings, i)
			ways := 0
			for length := 1; length <= maxLen; length++ {
				ways = add(ways, dp[i-length][j-length*endings])
			}
			dp[i][j] = add(dp[i][j], mul(25-w, ways))
		}
		for j := range k + 1 {
			pref[j] = add(pref[j], dp[i][j])
		}
	}

	ans := 0
	if k == 0 {
		ans = 1
	}
	for i := 1; i <= n; i++ {
		ans = add(ans, dp[i][k])
	}
	return ans
}
