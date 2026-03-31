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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(a, b)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(a []int, b []int) int {

	n := len(a)
	// dp[i][0] 表示不交换a[i]和b[i]时，满足条件的ways
	// dp[i][1] 表示交换时的结果
	dp := []int{1, 1}

	play := func(x int, y int, d int) (int, int) {
		if d == 1 {
			x, y = y, x
		}
		return x, y
	}

	for i := 1; i < n; i++ {
		ndp := []int{0, 0}
		for d := range 2 {
			x, y := play(a[i-1], b[i-1], d)
			for nd := range 2 {
				nx, ny := play(a[i], b[i], nd)
				if x <= nx && y <= ny {
					ndp[nd] = add(ndp[nd], dp[d])
				}
			}
		}
		dp = ndp
	}

	return add(dp[0], dp[1])
}
