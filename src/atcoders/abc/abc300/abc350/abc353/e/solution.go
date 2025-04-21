package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) int {
	readString(reader)
	s := readString(reader)
	x := strings.Split(s, " ")
	return solve(x)
}

func solve(s []string) int {
	// 先要将s排序？
	sort.Strings(s)

	n := len(s)
	a := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		a[i] = lcp(s[i], s[i+1])
	}

	stack := make([]int, n)
	var top int

	lf := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		for top > 0 && a[stack[top-1]] >= a[i] {
			top--
		}
		if top == 0 {
			lf[i] = -1
		} else {
			lf[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}

	top = 0

	var res int

	for i := n - 2; i >= 0; i-- {
		for top > 0 && a[stack[top-1]] > a[i] {
			top--
		}
		r := n - 1
		if top > 0 {
			r = stack[top-1]
		}
		res += a[i] * (i - lf[i]) * (r - i)
		stack[top] = i
		top++
	}
	return res
}

func lcp(a, b string) int {
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return min(len(a), len(b))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
