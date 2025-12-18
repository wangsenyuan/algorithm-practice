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
	var k int
	fmt.Fscan(reader, &k)
	c := make([]int, k)
	for i := range k {
		fmt.Fscan(reader, &c[i])
	}
	return solve(n, m, c)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

const inf = 1 << 60

func solve(n int, m int, c []int) int {
	if m == 1 {
		return 1
	}

	N := 1 << n
	dp := make([][]int, N)

	for i := range N {
		dp[i] = make([]int, n+1)
		for j := range n + 1 {
			dp[i][j] = -1
		}
	}

	good := make([]bool, n)
	for _, i := range c {
		good[i-1] = true
	}

	var f func(state int, l int) int
	f = func(state int, l int) (res int) {
		if l == 1 {
			return (state & 1) + 1
		}

		if dp[state][l] != -1 {
			return dp[state][l]
		}

		defer func() {
			dp[state][l] = res
		}()

		res = f(state>>1, l-1)

		fn := func(a, b int) int {
			return min(a, b)
		}
		if (n-l)&1 == 0 {
			// alice, remove 0, bob remove 1
			fn = func(a, b int) int {
				return max(a, b)
			}
		}

		for i := range l {
			if i > 0 && good[i] {
				newState := state>>(i+1)<<i | (state & (1<<i - 1))
				newState &= (1<<(l-1) - 1)
				res = fn(res, f(newState, l-1))
			}
		}

		return
	}

	var res int

	for i := range N {
		res = add(res, f(i, n))
	}

	return res
}
