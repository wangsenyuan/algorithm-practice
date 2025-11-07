package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, a []int) []int {
	ans := make([]int, n)
	for i := range n {
		ans[i] = n
	}
	ans[0] = 0

	que := make([]int, n)
	var head, tail int
	que[head] = 0
	head++
	for tail < head {
		u := que[tail]
		tail++
		v := a[u] - 1
		if u < v && ans[v] > ans[u]+1 {
			ans[v] = ans[u] + 1
			que[head] = v
			head++
		}

		if u+1 < n && ans[u+1] > ans[u]+1 {
			ans[u+1] = ans[u] + 1
			que[head] = u + 1
			head++
		}
		if u-1 > 0 && ans[u-1] > ans[u]+1 {
			ans[u-1] = ans[u] + 1
			que[head] = u - 1
			head++
		}
	}

	return ans
}
