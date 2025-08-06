package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func process(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(k, a)
}

func solve(k int, a []int) []int {
	// 按顺序处理就好了
	var res []int
	var sum int
	n := len(a)
	for i := 1; i < n; i++ {
		cnt := i - len(res)
		score := sum - cnt*(n-(i+1))*a[i]
		if score < k {
			res = append(res, i+1)
		} else {
			sum += a[i] * cnt
		}
	}
	return res
}
