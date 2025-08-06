package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	fmt.Println(ans[0], ans[1])
}

func process(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, k, a)
}

func solve(n int, k int, a []int) []int {
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	best := make([]int, n+1)
	best[n] = n
	for i := n - 1; i >= 0; i-- {
		best[i] = i
		if best[i+1]+k <= n && sum[i+k]-sum[i] < sum[best[i+1]+k]-sum[best[i+1]] {
			best[i] = best[i+1]
		}
	}
	ans := make([]int, 2)
	ans[0] = 0
	ans[1] = best[k]
	tmp := sum[k] - sum[0] + sum[best[k]+k] - sum[best[k]]
	for i := 1; i+2*k <= n; i++ {
		tmp1 := sum[i+k] - sum[i] + sum[best[i+k]+k] - sum[best[i+k]]
		if tmp1 > tmp {
			tmp = tmp1
			ans[0] = i
			ans[1] = best[i+k]
		}
	}

	ans[0]++
	ans[1]++
	return ans
}
