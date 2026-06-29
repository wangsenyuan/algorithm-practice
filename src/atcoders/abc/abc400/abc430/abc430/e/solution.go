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
	var a, b string
	fmt.Fscan(reader, &a)
	fmt.Fscan(reader, &b)
	return solve(a, b)
}

func solve(a, b string) int {
	n := len(a)
	a += a
	p := kmp(b)

	var j int
	for i := range 2 * n {
		for j > 0 && a[i%n] != b[j] {
			j = p[j-1]
		}
		if a[i%n] == b[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}

	return -1
}

func kmp(s string) []int {
	n := len(s)
	p := make([]int, n)
	for i := 1; i < n; i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
	return p
}
