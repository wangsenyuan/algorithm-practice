package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, k, X int

	fmt.Fscan(reader, &n, &k, &X)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}

	return solve(k, X, a)
}

func solve(k int, X int, a []int) int {
	slices.Sort(a)

	n := len(a)
	sum := make([]int, n+1)

	for i, v := range a {
		sum[i+1] = sum[i] + v
	}

	if sum[n] < X {
		return -1
	}

	// 假设一共喝了w杯水, 剩余 n - w 杯
	// 如果这n - w杯都是水, n - w < k, 那么 肯定有 v = k - (n - w)杯是sake
	// 那么最差情况下,就是最少的这v杯是sake
	// 如果这v杯水的 sum >= X, 那么答案可以是w
	for w := n - k + 1; w <= n; w++ {
		if sum[k]-sum[n-w] >= X {
			return w
		}
	}

	return -1
}
