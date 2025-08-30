package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	res := solve(s)
	fmt.Println(res)
}

func solve(s string) int {
	var dots []int
	var ats []int
	n := len(s)
	for i := range n {
		if s[i] == '.' {
			dots = append(dots, i)
		}
		if s[i] == '@' {
			ats = append(ats, i)
		}
	}

	// next pos of non-letters
	next := make([]int, n+1)
	next[n] = n
	for i := n - 1; i >= 0; i-- {
		if s[i] >= 'a' && s[i] <= 'z' {
			next[i] = next[i+1]
		} else {
			next[i] = i
		}
	}

	sum := make([]int, n+1)
	for i := range n {
		sum[i+1] = sum[i]
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= '0' && s[i] <= '9' {
			sum[i+1]++
		}
	}

	var ans int

	var j int
	for i := 0; i < len(ats); i++ {
		// 要找到i前面的.或者@
		var l int
		if i > 0 {
			l = ats[i-1] + 1
		}
		for j < len(dots) && dots[j] < ats[i] {
			l = max(l, dots[j]+1)
			j++
		}
		// dots[j] > ats[i]
		if j == len(dots) {
			// not valid
			break
		}
		if ats[i]+1 == dots[j] || next[dots[j]+1] == dots[j]+1 || sum[dots[j]]-sum[ats[i]] != dots[j]-ats[i]-1 {
			// dots[j] + 1 不是一个letter
			continue
		}
		// 中间必须是letters或者是numbers
		r := next[dots[j]+1]
		for l < ats[i] {
			if s[l] >= 'a' && s[l] <= 'z' {
				ans += r - dots[j] - 1
			}
			l++
		}
	}
	return ans
}
