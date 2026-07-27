package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	ops := make([][]int, q)
	for i := range q {
		ops[i] = make([]int, 2)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1])
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) []int {
	cnt := make([]int, n+1)
	var que []int
	for i := 1; i <= n; i++ {
		que = append(que, i)
		cnt[i] = 1
	}

	ans := make([]int, len(ops))

	for i, cur := range ops {
		x, y := cur[0], cur[1]
		var sum int
		for len(que) > 0 && que[0] <= x {
			sum += cnt[que[0]]
			cnt[y] += cnt[que[0]]
			que = que[1:]
		}
		ans[i] = sum
	}

	return ans
}
