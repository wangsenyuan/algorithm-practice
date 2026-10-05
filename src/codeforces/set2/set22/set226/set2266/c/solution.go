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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	var s string
	fmt.Fscan(reader, &n, &s)
	_ = n
	return solve(s)
}

func solve(s string) int {
	n := len(s)
	if s[0] == '1' {
		// 只能使用OR操作
		var ans int
		for i := range n {
			if s[i] == '0' {
				ans++
			}
		}
		return ans
	}
	firstOne := -1
	for i := range n {
		if s[i] == '1' {
			firstOne = i
			break
		}
	}
	if firstOne == -1 {
		return 0
	}
	pref := make([]int, n)
	for i := firstOne; i < n; i++ {
		pref[i] = pref[i-1]
		if s[i] == '1' {
			pref[i]++
		}
	}
	// 全部变成0
	best := pref[n-1]
	var suf int
	for i := n - 1; i >= firstOne; i-- {
		if s[i] == '0' {
			suf++
		}
		// 后面全部变成1
		best = min(best, pref[i-1] + suf)
	}

	return best
}
