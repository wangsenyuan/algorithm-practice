package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n, k int
	fmt.Fscan(reader, &n, &k)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	return solve(n, k, a)
}

const N = 200010

var lpf [N]int

func init() {
	var primes []int
	for i := 2; i < N; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= N {
				break
			}
			lpf[i*j] = j
			if i%j == 0 {
				break
			}
		}
	}
}

func solve(n, k int, a []int) int64 {
	dp := make([]int, n+1)
	// dp[0...k] = 0
	for i := k + 1; i <= n; i++ {
		num := i
		dp[i] = i
		for num > 1 {
			// 选择 p = lpf[num]
			dp[i] = min(dp[i], lpf[num]*dp[i/lpf[num]]+1)
			num /= lpf[num]
		}
	}
	var ans int
	for _, num := range a {
		ans += dp[num]
	}
	return int64(ans)
}
