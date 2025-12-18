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
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) bool {
	a := readString(reader)
	b := readString(reader)
	return solve(a, b)
}

func solve(a string, b string) bool {
	count := func(s string) int {
		var res int
		for i := range len(s) {
			res += int(s[i] - '0')
		}
		return res
	}

	c1 := count(a)
	c2 := count(b)
	if c1&1 == 1 {
		// 奇数可以变成偶数
		c1++
	}
	return c2 <= c1
}
