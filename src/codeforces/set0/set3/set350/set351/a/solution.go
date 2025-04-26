package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%.3f\n", process(reader))
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func process(reader *bufio.Reader) float64 {
	n := readNum(reader) * 2
	a := make([]int, n)
	var tmp int
	for i := range n {
		fmt.Fscanf(reader, "%d.%d", &tmp, &a[i])
	}
	return solve(a)
}

const inf = 1 << 60

func solve(a []int) float64 {
	n := len(a)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n/2+1)
		for j := range dp[i] {
			dp[i][j] = -inf
		}
	}
	var dfs func(i int, j int) int

	dfs = func(i int, j int) int {
		if j > n/2 {
			return inf
		}
		if i == n {
			if j == n/2 {
				return 0
			}
			return inf
		}
		if dp[i][j] != -inf {
			return dp[i][j]
		}
		tmp1 := dfs(i+1, j+1) + a[i]
		// 求余是必要的， a[i] = 0
		tmp2 := dfs(i+1, j) - (1000-a[i])%1000
		if abs(tmp1) < abs(tmp2) {
			dp[i][j] = tmp1
		} else {
			dp[i][j] = tmp2
		}
		return dp[i][j]
	}

	ans := dfs(0, 0)
	return float64(abs(ans)) / 1000.0
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
