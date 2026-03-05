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

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return (a * b) % mod
}

func drive(reader *bufio.Reader) int {
	var n, h int
	fmt.Fscan(reader, &n, &h)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, h)
}

func solve(a []int, h int) int {
	n := len(a)
	if n == 1 && (a[0] == h || a[0]+1 == h) {
		return 1
	}
	// dp[i] 表示到目前为止有多少个还没有结束的区间
	dp := make([]int, n+1)

	if h == a[0] {
		dp[0] = 1
		// 这里不能有开始的区间
	} else {
		// w == h[0] + 1
		dp[1] = 1
		// 第一个区间 [0, 0]
		dp[0] = 1
	}
	for i := 1; i < n; i++ {
		if h < a[i] {
			// 无法满足最后一样高的条件
			return 0
		}
		// 这里需要增加d次
		d := h - a[i]
		if d > i+1 {
			return 0
		}
		if d > 0 {
			// 这里可以不开始新区间，或者开始，或者开始/结束
			// 不开始的话, dp[d] 保留
			// 开始的话， dp[d] += dp[d-1], (只能是前面有d-1个，然后开始一个新区间)
			// 开始/结束 dp[d-1] *= d (因为有d个位置可以选择， 包括自己)
			x := dp[d]
			y := dp[d-1]
			for j := range n + 1 {
				dp[j] = 0
			}
			dp[d] = add(x, y)
			dp[d-1] = mul(d, dp[d])
		} else {
			// d = 0, 必须全部开始区间结束掉了，只能保留dp[0]
			for j := 1; j <= n; j++ {
				dp[j] = 0
			}
		}
	}
	return dp[0]
}
