package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	ok, res := solve(s)
	if ok {
		fmt.Println("YES")
		fmt.Println(res)
	} else {
		fmt.Println("NO")
	}
}

func solve(s string) (bool, string) {
	for i := range len(s) {
		if s[i] == '0' {
			return true, "0"
		}
	}

	// 不包括0
	n := len(s)

	dp := make([]int, n)

	for i := range n {
		x := int(s[i] - '0')
		if i > 0 {
			dp[i] = dp[i-1]
			for j := range 8 {
				if (dp[i-1]>>j)&1 == 1 {
					dp[i] |= 1 << ((j*10 + x) % 8)
				}
			}
		}
		dp[i] |= 1 << (x % 8)
	}

	if dp[n-1]&1 == 0 {
		return false, ""
	}
	var buf []byte

	w := 0

	for i := n - 1; i >= 0 && w >= 0; i-- {
		if i > 0 && (dp[i-1]>>w)&1 == 1 {
			continue
		}
		buf = append(buf, s[i])
		if i == 0 {
			break
		}
		x := int(s[i] - '0')
		nw := -1
		for j := range 8 {
			if (dp[i-1]>>j)&1 == 1 && (j*10+x)%8 == w {
				nw = j
				break
			}
		}
		w = nw
	}

	// w == 0

	slices.Reverse(buf)

	return true, string(buf)
}
