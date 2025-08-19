package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	ops := make([][]int, q)
	for i := range q {
		ops[i] = make([]int, 3)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1], &ops[i][2])
	}
	return solve(n, ops)
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(n int, ops [][]int) []int {

	open := make([][]int, n+1)
	close := make([][]int, n+1)
	for i, cur := range ops {
		l, r := cur[0], cur[1]
		open[l] = append(open[l], i)
		close[r] = append(close[r], i)
	}

	ok := make([]bool, n+1)
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for _, j := range open[i] {
			x := ops[j][2]
			for k := n; k >= x; k-- {
				dp[k] = add(dp[k], dp[k-x])
			}
		}
		for k := 1; k <= n; k++ {
			if dp[k] > 0 {
				ok[k] = true
			}
		}
		for _, j := range close[i] {
			x := ops[j][2]
			for k := x; k <= n; k++ {
				dp[k] = sub(dp[k], dp[k-x])
			}
		}

	}

	var res []int
	for k := 1; k <= n; k++ {
		if ok[k] {
			res = append(res, k)
		}
	}
	return res
}
