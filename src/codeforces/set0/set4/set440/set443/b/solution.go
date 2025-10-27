package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	var k int
	fmt.Fscan(reader, &k)
	return solve(s, k)
}

func solve(s string, k int) int {
	n := len(s)

	if k >= n {
		return (n + k) / 2 * 2
	}

	// j + m < n + k and (j + m >= n or s[j] == s[j+m])
	check := func(m int) bool {
		if m <= k {
			// m <= n also holds
			return true
		}
		for l := 0; l + 2 * m <= n + k; l++ {
			ok := true
			for i := 0; i < m && l + i < n; i++ {
				if l + i + m < n && s[l+i] != s[l+i+m] {
					ok = false
					break
				}
			}
			if ok {
				return true
			}
		}

		return false
	}

	m := (n + k) / 2
	for m > 0 && !check(m) {
		m--
	}
	return 2 * m
}
