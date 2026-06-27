package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	var s, z string
	fmt.Fscan(reader, &s, &z)
	return solve(n, k, s, z)
}

func solve(n, k int, s, z string) int64 {
	cnt := make([]int64, 4)
	for i := range n {
		u := int(s[i] - '0')
		v := int(z[i] - '0')
		cnt[u<<1|v]++
	}

	score := calc(k)
	var res int64
	for i := 0; i < 4; i++ {
		for j := i; j < 4; j++ {
			var ways int64
			if i == j {
				ways = cnt[i] * (cnt[i] - 1) / 2
			} else {
				ways = cnt[i] * cnt[j]
			}
			a := (i >> 1) ^ (j >> 1)
			b := (i & 1) ^ (j & 1)
			res += ways * score[a][b]
		}
	}
	return res
}

func calc(k int) [][]int64 {
	dp := make([][]int64, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int64, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = int64(i + j)
		}
	}
	for ; k > 0; k-- {
		fp := make([][]int64, 2)
		for i := 0; i < 2; i++ {
			fp[i] = make([]int64, 2)
			for j := 0; j < 2; j++ {
				x := i ^ j
				fp[i][j] = dp[i][x] + dp[x][j] - int64(x)
			}
		}
		dp = fp
	}
	return dp
}
