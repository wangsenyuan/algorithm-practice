package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Println(drive(reader))
	}
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	return solve(s)
}

func solve(s string) int {
	var cnt int
	// )( 有多少个
	n := len(s)
	var phase int
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			phase++
		}
		if phase > 0 && s[i] == '(' {
			cnt++
		}
		if cnt == 2 {
			return n - 2
		}
	}
	return -1
}
