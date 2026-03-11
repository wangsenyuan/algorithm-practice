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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) string {
	n := len(a)

	var sum int
	for _, v := range a {
		sum += v
	}
	if sum == 0 {
		return "BitAryo"
	}

	if n == 1 {
		return "BitLGM"
	}
	if n == 2 {
		return solve2(a)
	}
	cnt := boolToInt(a[0] > 0) + boolToInt(a[1] > 0) + boolToInt(a[2] > 0)

	if cnt == 1 {
		return "BitLGM"
	}
	if cnt == 2 {
		w := a[0] ^ a[1] ^ a[2]
		if w == 0 {
			return "BitAryo"
		}
	}

	if a[0]^a[1]^a[2] == 0 {
		return "BitAryo"
	}
	// 三个的时候，会超时
	return "BitLGM"
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func solve2(a []int) string {
	dp := make([][]int, a[0]+1)
	for i := range a[0] + 1 {
		dp[i] = make([]int, a[1]+1)
		for j := range a[1] + 1 {
			dp[i][j] = -1
		}
	}

	var f func(u int, v int) int
	f = func(u int, v int) (res int) {
		cnt := boolToInt(u > 0) + boolToInt(v > 0)
		if cnt == 1 {
			return 1
		}
		// cnt == 2
		if u == v {
			return 1
		}
		if dp[u][v] != -1 {
			return dp[u][v]
		}
		defer func() {
			dp[u][v] = res
		}()
		for x := 1; x <= u; x++ {
			if f(u-x, v) == 0 {
				return 1
			}
		}
		for x := 1; x <= v; x++ {
			if f(u, v-x) == 0 {
				return 1
			}
		}
		x1 := min(u, v)
		for x := 1; x <= x1; x++ {
			if f(u-x, v-x) == 0 {
				return 1
			}
		}
		return
	}

	res := f(a[0], a[1])
	if res == 1 {
		return "BitLGM"
	}
	return "BitAryo"
}
