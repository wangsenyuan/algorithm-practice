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
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Fprintln(writer, s[1:len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	return solve(k, a, b)
}


func solve(k int, a, b []int) []int {
	n := len(a)
	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = a[i]/b[i] + suf[i+1]
	}
	if suf[0] < k {
		// 完不成任务
		return make([]int, n)
	}
	ans := make([]int, n)

	var tot int
	for i := range n {
		// 假设c[i] = x,
		// 那么 x + tot + suf[i+1] >= k
		ans[i] = max(0, k-suf[i+1]-tot)
		tot += ans[i]
	}

	return ans
}
