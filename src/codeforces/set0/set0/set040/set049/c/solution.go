package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	ans := solve(n)
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int) []int {
	if n == 1 {
		return []int{1}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = i + 2
	}
	ans[n-1] = 1
	return ans
}
