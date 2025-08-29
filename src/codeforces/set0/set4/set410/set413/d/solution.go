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

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	return solve(k, a)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

const H = 12

func solve(k int, a []int) int {
	n := len(a)
	T := 1 << H
	dp := make([]int, T*2)
	if a[0] != 0 {
		dp[1<<(a[0]/2)]++
	} else {
		dp[1<<1]++
		dp[1<<2]++
	}

	// 除了mask，还必须知道这个mask在历史上，是否产生过超过k的位

	get := func(mask int, i int) int {
		// 比i低的状态要清理掉
		if mask&(1<<i-1) == 0 {
			for j := i; j < H; j++ {
				// 进位，知道遇到空的位置
				if (mask>>j)&1 == 0 {
					mask ^= 1 << j
					break
				}
				mask ^= 1 << j
			}
		} else {
			mask = (mask & (1 << H)) | 1<<i
		}

		if (mask>>k)&1 == 1 {
			// 一旦得到了位k
			mask |= 1 << H
		}

		return mask
	}

	ndp := make([]int, T*2)
	// var res int
	for i := 1; i < n; i++ {
		for mask := range T * 2 {
			if dp[mask] == 0 {
				continue
			}
			if a[i] == 0 {
				next := get(mask, 1)
				ndp[next] = add(ndp[next], dp[mask])
				next = get(mask, 2)
				ndp[next] = add(ndp[next], dp[mask])
			} else {
				next := get(mask, a[i]/2)
				ndp[next] = add(ndp[next], dp[mask])
			}
		}
		copy(dp, ndp)
		clear(ndp)
	}

	var ans int
	for mask := T; mask < 2*T; mask++ {
		// 如果mask比k大的位有甚至
		ans = add(ans, dp[mask])
	}
	return ans
}
