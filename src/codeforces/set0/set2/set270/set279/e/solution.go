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

const inf = 1 << 60

func solve(s string) int {
	n := len(s)
	t := make([]byte, n+1)
	copy(t[1:], s)
	for i := 1; i <= n; i++ {
		if t[i] == '0' {
			t[i] = '1'
		} else {
			t[i] = '0'
		}
	}
	// t + 1
	p := n
	for p > 0 && t[p] == '1' {
		t[p] = '0'
		p--
	}
	if p == 0 {
		t[0] = '1'
		s = "0" + s
	} else {
		t[p] = '1'
		t = t[1:]
	}
	n = len(s)

	v := []string{s, string(t)}

	dp := make([]int, 2)
	if s[n-1] == '1' {
		dp[0] = 1
	}
	if t[n-1] == '1' {
		dp[1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		fp := []int{inf, inf}
		for b := range 2 {
			if v[b][i] == '0' {
				fp[b] = dp[b]
			} else {
				fp[b] = min(dp[b], dp[b^1]) + 1
			}
		}
		copy(dp, fp)
	}

	return dp[0]
}
