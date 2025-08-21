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
	s := fmt.Sprintf("%v", ans)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int, k int) []int {
	k = min(k, n-k)

	ans := make([]int, n)
	var cycles int
	sum := 1
	var cur int
	for i := range n {
		to := (cur + k) % n
		if to < cur && to != 0 {
			cycles++
		}
		sum += cycles + 1
		if to < cur {
			cycles++
		}
		cur = to
		ans[i] = sum
	}
	return ans
}
