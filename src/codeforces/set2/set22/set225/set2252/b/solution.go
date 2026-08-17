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
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		res := solve(s)
		fmt.Fprintln(writer, res)
	}
}

func solve(s string) int {
	// 要将s变成0101 或者1010, 删除最小的个数
	// 假设删除了x个1, y个0, 且 x - y = 1
	cnt := make([]int, 2)
	n := len(s)
	for i := range n {
		cnt[int(s[i]-'0')]++
	}

	play := func(pat string) int {
		// pat = 01 or 10
		var w int
		sum := make([]int, 2)
		res := -1
		for i := 0; i < n; i++ {
			if s[i] == pat[w] {
				w ^= 1
				sum[int(s[i]-'0')]++
				// 这个是需要删除的部分
				x0 := cnt[0] - sum[0]
				x1 := cnt[1] - sum[1]
				if abs(x0-x1) <= 1 {
					res = sum[0] + sum[1]
				}
			}
		}

		return res
	}
	ans := max(play("01"), play("10"))
	if ans == -1 {
		return -1
	}
	return n - ans
}

func abs(num int) int {
	return max(num, -num)
}
