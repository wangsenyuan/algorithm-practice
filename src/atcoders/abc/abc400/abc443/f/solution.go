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

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	return solve(n)
}

func solve(n int) string {
	// n 的倍数
	dp := make([]int, n*10)
	fp := make([]int, n*10)
	for i := range n * 10 {
		dp[i] = -1
		fp[i] = -1
	}

	var que []int
	for i := 1; i <= 9; i++ {
		dp[(i%n)*10+i] = 1
		que = append(que, (i%n)*10+i)
	}

	best := -1

	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		w, d := cur/10, cur%10
		if w == 0 {
			best = cur
			break
		}
		for d1 := d; d1 <= 9; d1++ {
			next := (w*10+d1)%n*10 + d1
			if dp[next] != -1 {
				break
			}
			dp[next] = dp[cur] + 1
			fp[next] = cur
			que = append(que, next)
		}
	}

	if best < 0 {
		return "-1"
	}
	var buf []byte

	for best > 0 {
		d := best % 10
		buf = append(buf, byte(d+'0'))
		best = fp[best]
	}
	slices.Reverse(buf)
	return string(buf)
}
