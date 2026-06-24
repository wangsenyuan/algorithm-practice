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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for range tc {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)
	var s string
	fmt.Fscan(reader, &s)
	return solve(n, s)
}

func solve(n int, s string) string {
	var cnt int
	for i := range n {
		if s[i] == 'W' {
			cnt++
		}
	}
	if cnt%2 != 0 {
		return "NO"
	}
	cnt /= 2
	for i := range cnt {
		if s[i] != 'W' {
			return "NO"
		}
	}
	for i := 2*n - cnt; i < 2*n; i++ {
		if s[i] != 'R' {
			return "NO"
		}
	}

	return "YES"
}
