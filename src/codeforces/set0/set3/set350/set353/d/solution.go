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
	for len(s) > 0 && s[0] == 'F' {
		s = s[1:]
	}

	if len(s) == 0 {
		return 0
	}
	// s[0] = 'M'
	var res int
	var cnt int
	var prev = -1
	var last = -1
	for i := 0; i < len(s); i++ {
		if s[i] == 'F' {
			if prev == -1 {
				// 第一个
				res = i
				last = i
			} else {
				if i-cnt > last {
					last = i - cnt
				} else {
					last++
				}
				res = max(res, last)
			}
			cnt++
			prev = i
		}
	}

	return res
}
