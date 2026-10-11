package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

func solve(s string) int {
	var suf int
	for i := range s {
		if isUpper(s[i]) {
			suf++
		}
	}
	ans := suf
	var pref int
	for i := range s {
		if isUpper(s[i]) {
			suf--
		}
		// 把前面所有的小写字母变成大写, 把后面所有的大写变成小写
		ans = min(ans, pref+suf)
		if !isUpper(s[i]) {
			pref++
		}
	}
	return ans
}

func isUpper(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}
