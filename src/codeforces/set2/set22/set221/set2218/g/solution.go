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
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(m, b)
}

const mod = 676767677

func mul(a, b int) int {
	return a * b % mod
}

func solve(m int, b []int) int {
	n := len(b)

	open := make([][]int, m)
	for i, v := range b {
		open[v] = append(open[v], i)
	}

	cnt := make([]int, m)
	cnt[0] = len(open[0])

	ans := 1

	for d := 1; d < m; d++ {
		for _, j := range open[d] {
			if (j == 0 || b[j-1] >= d) && (j == n-1 || b[j+1] >= d) {
				// invalid, 两边还没有坐下
				return 0
			}
			lo := 1
			if d > 1 && (j > 0 && b[j-1] < d-1) || (j < n-1 && b[j+1] < d-1) {
				lo = cnt[d-2] + 1
			}
			ans = mul(ans, cnt[d-1]-lo+1)
		}
		cnt[d] = cnt[d-1] + len(open[d])
	}

	return ans
}
