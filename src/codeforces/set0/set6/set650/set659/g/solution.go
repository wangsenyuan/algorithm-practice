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
	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	h := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &h[i])
	}
	return solve(h)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(h []int) int {
	n := len(h)
	ans := h[0] - 1
	if n == 1 {
		return ans
	}

	dp := min(h[0], h[1]) - 1

	for i := 1; i < n; i++ {
		ans = add(ans, h[i]-1)
		ans = add(ans, mul(dp, min(h[i-1], h[i])-1))
		if i+1 < n {
			dp = mul(dp, min(h[i-1], h[i], h[i+1])-1)
			dp = add(dp, min(h[i], h[i+1])-1)
		}
	}

	return ans
}
