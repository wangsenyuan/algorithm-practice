package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) int {
	n := len(s)
	if s[n-1] == 'R' {
		s = revert(s)
	}

	// 在右边放
	play := func(obs int) bool {
		l := 0
		var x int
		for i := 0; i < n-1; i++ {
			if s[i] == 'L' {
				x--
				l = min(l, x)
			} else {
				if x+1 < obs {
					x++
				}
				// else stay no change
			}
		}
		// 最后一个肯定是L
		return x == l
	}

	if play(n + 2) {
		// 没有任何限制时
		return 1
	}

	l, r := 1, n+1
	for l < r {
		mid := (l + r) / 2
		if play(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return r - 1
}

func revert(s string) string {
	buf := []byte(s)
	for i := range len(buf) {
		if buf[i] == 'L' {
			buf[i] = 'R'
		} else {
			buf[i] = 'L'
		}
	}
	return string(buf)
}
