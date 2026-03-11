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
	if len(a) != len(b) {
		return false
	}
	if a == b {
		return true
	}

	get := func(s string) int {
		var res int
		for _, c := range s {
			res += int(c - '0')
		}
		return res
	}

	x := get(a)
	y := get(b)
	// 如果a全部是0，那么a没法变化
	// 如果b全部是0，虽然a可以变，但是a的1的个数，至少会是1
	if x == 0 || y == 0 {
		return false
	}

	return true
}
