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
	var n, k, p, q int
	fmt.Fscan(reader, &n, &k, &p, &q)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, k, p, q, a)
}

func solve(n, k, p, q int, a []int) int64 {
	fp := make([]int, n+1)
	dp1 := make([]int, n+1)
	dp2 := make([]int, n+2)
	for i, v := range a {
		fp[i+1] = fp[i] + min(v%p, v%q%p)
		dp1[i+1] = dp1[i] + v%p
		dp2[i+1] = dp2[i] + v%q%p
	}

	best := min(dp1[n], dp2[n])
	var suf int
	for i := n - 1; i >= k-1; i-- {
		j := i - k + 1
		tmp := fp[j] + suf + min(dp1[i+1]-dp1[j], dp2[i+1]-dp2[j])
		best = min(best, tmp)
		suf += min(a[i]%p, a[i]%q%p)
	}

	return int64(best)
}
