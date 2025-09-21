package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, a, b, c int
	fmt.Fscan(reader, &n, &a, &b, &c)
	return solve(n, a, b, c)
}

func solve(n int, a int, b int, c int) int {
	if a <= b-c || n < b {
		return n / a
	}
	// a > b - c
	// n - m * (b - c) >= b
	// m * (b - c) <= n - b
	m := (n - b) / (b - c)
	res := m
	n -= m * (b - c)
	if n >= b {
		res++
		n -= b - c
	}
	res += n / a
	return res
}
