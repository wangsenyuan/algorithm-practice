package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	readString(reader)
	s := readString(reader)
	return solve(s)
}

const mod = 1_000_000_007

func add(a, b int32) int32 {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(s string) int {
	n := len(s)

	var sum, M int
	for sum < n {
		if sum+bits.Len(uint(M+1)) > n {
			break
		}
		M++
		sum += bits.Len(uint(M))
	}

	dp := make([][]int32, n+1)
	for i := range n + 1 {
		dp[i] = make([]int32, 1<<M)
	}

	var ans int32

	for i := range n {
		var num int
		for j := i; j < n; j++ {
			num = num*2 + int(s[j]-'0')
			if num == 0 {
				continue
			}
			if num > M {
				break
			}
			dp[j+1][1<<(num-1)] = add(dp[j+1][1<<(num-1)], 1)
			for mask := 1; mask < 1<<M; mask++ {
				next := mask | (1 << (num - 1))
				dp[j+1][next] = add(dp[j+1][next], dp[i][mask])
			}
		}
		for mask := 2; mask <= 1<<M; mask *= 2 {
			ans = add(ans, dp[i+1][mask-1])
		}
	}
	return int(ans)
}
