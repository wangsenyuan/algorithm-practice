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

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	ans := make([]int, t)
	for i := range t {
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		ans[i] = solve(n, s)
	}
	return ans
}

func solve(n int, s string) int {
	suf := make([]int, n+1)
	cnt := make([]int, 2)

	for i := n - 1; i >= 0; i-- {
		d := int(s[i] - '0')
		suf[i] = cnt[1-d] + 2*cnt[d]
		cnt[d]++
	}

	clear(cnt)

	best := suf[0]

	for i := 0; i < n; {
		j := i
		for i < n && s[i] == s[j] {
			i++
		}
		d := int(s[j] - '0')
		best = min(best, cnt[1-d]+2*cnt[d]+suf[i-1])
		cnt[d] += i - j
	}

	return best
}
