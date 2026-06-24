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
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	var s string
	fmt.Fscan(reader, &s)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(n, x, y, s, p)
}

func solve(n int, x int, y int, s string, p []int) string {
	if x+y < sum(p) {
		return "NO"
	}
	sum := make([]int, 2)
	var flag int
	for i := range n {
		win := p[i]/2 + 1
		if s[i] == '0' {
			flag |= 1
			sum[0] += win
		} else {
			flag |= 2
			sum[1] += win
		}
	}
	if flag == 3 {
		if x < sum[0] || y < sum[1] {
			return "NO"
		}

		return "YES"
	}

	if s[0] == '0' && (x >= sum[0] && x >= y+n) {
		return "YES"
	}
	if s[0] == '1' && (y >= sum[1] && y >= x+n) {
		return "YES"
	}
	return "NO"
}

func sum(arr []int) int {
	var res int
	for _, v := range arr {
		res += v
	}
	return res
}
