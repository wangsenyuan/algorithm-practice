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
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for j := range m {
		fmt.Fscan(reader, &b[j])
	}
	return solve(a, b)
}

func solve(a []int, b []int) int {
	m := len(b)
	if m == 1 {
		return len(a)
	}
	diff := make([]int, m-1)
	for i := 0; i+1 < m; i++ {
		diff[i] = b[i+1] - b[i]
	}
	next := kmp(diff)

	n := len(a)
	var ans int

	var j int
	for i := 0; i+1 < n; i++ {
		cur := a[i+1] - a[i]
		for j > 0 && cur != diff[j] {
			j = next[j-1]
		}
		if cur == diff[j] {
			j++
		}
		if j == m-1 {
			ans++
			j = next[j-1]
		}
	}
	return ans
}

func kmp(p []int) []int {
	n := len(p)
	next := make([]int, n)
	for i := 1; i < n; i++ {
		j := next[i-1]
		for j > 0 && p[i] != p[j] {
			j = next[j-1]
		}
		if p[i] == p[j] {
			j++
		}
		next[i] = j
	}
	return next
}
