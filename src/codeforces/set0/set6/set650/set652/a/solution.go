package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var h1, h2, a, b int
	fmt.Fscan(reader, &h1, &h2, &a, &b)
	return solve(h1, h2, a, b)
}

func solve(h1 int, h2 int, a int, b int) int {
	if h2-h1 <= 8*a {
		return 0
	}
	if a <= b {
		return -1
	}
	// 第 0 天已经爬了 8 小时。之后每个白天 12 小时、夜里 12 小时，净高度是 12*(a-b)
	need := h2 - h1 - 8*a
	day := 12 * (a - b)
	return (need + day - 1) / day
}
