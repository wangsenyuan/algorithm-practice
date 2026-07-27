package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, v := range drive(reader) {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []string {
	var t int
	fmt.Fscan(reader, &t)
	ans := make([]string, t)
	for i := range t {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		ans[i] = solve(a)
	}
	return ans
}

func solve(a []int) string {
	var cnt int
	for _, v := range a {
		if v == 1 {
			cnt++
		}
	}
	w := slices.Max(a)
	if w > 1 {
		cnt++
	}
	if cnt&1 == 0 {
		return "Bob"
	}
	return "Alice"
}
