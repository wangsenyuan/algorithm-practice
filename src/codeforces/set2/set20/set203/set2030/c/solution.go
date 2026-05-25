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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	var s string
	fmt.Fscan(reader, &n, &s)
	return solve(s)
}

func solve(s string) string {
	n := len(s)
	if s[0] == '1' || s[n-1] == '1' {
		return "YES"
	}
	for i := 0; i+1 < n; i++ {
		if s[i] == '1' && s[i+1] == '1' {
			return "YES"
		}
	}

	return "NO"
}
