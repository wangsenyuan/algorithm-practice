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

func drive(reader *bufio.Reader) []string {
	var t int
	fmt.Fscan(reader, &t)
	res := make([]string, t)
	for i := range t {
		var n int
		var s string
		fmt.Fscan(reader, &n, &s)
		_ = n
		res[i] = solve(s)
	}
	return res
}

func solve(s string) string {
	var left, right = 0, len(s) - 1
	for left < len(s) && s[left] == '0' {
		left++
	}
	for right >= 0 && s[right] == '1' {
		right--
	}
	if left > right {
		return "Bob"
	}
	cnt := 1
	for i := left; i < right; i++ {
		if s[i+1] == s[i] {
			cnt++
		} else {
			if cnt%2 == 1 {
				return "Alice"
			}
			cnt = 1
		}
	}
	if cnt%2 == 1 {
		return "Alice"
	}
	return "Bob"
}
