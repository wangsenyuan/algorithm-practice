package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)
	ans := solve(n, k)
	if len(ans) == 0 {
		fmt.Println(-1)
		return
	}
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int, k int) []int {
	if n/3 < k || k == 1 {
		return nil
	}

	// 1, 2, 3, 4, 5, 6
	ans := make([]int, n)

	for i := 0; i < 3*k; i++ {
		ans[i] = i/3 + 1
	}

	var pos int
	if k%2 == 1 {
		// 1, 2, 3, 4, 5, 6, 7, 8, 9
		// 1, 2, 4, 3, 5, 8, 6, 7, 9
		ans[2], ans[3] = ans[3], ans[2]
		ans[5], ans[7] = ans[7], ans[5]
		pos = 9
	}

	// 偶数个分组
	for pos < 3*k {
		ans[pos+2], ans[pos+3] = ans[pos+3], ans[pos+2]
		pos += 6
	}

	for i := pos; i < n; i++ {
		ans[i] = 1
	}

	return ans
}
