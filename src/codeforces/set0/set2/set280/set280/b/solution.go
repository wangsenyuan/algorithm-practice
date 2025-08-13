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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)

	stack := make([]int, n)
	var top int

	var best int

	for i := 0; i < n; i++ {
		for top > 0 && stack[top-1] < a[i] {
			x := stack[top-1]
			best = max(best, x^a[i])
			top--
		}
		stack[top] = a[i]
		top++
	}

	top = 0
	for i := n - 1; i >= 0; i-- {
		for top > 0 && stack[top-1] < a[i] {
			x := stack[top-1]
			best = max(best, x^a[i])
			top--
		}
		stack[top] = a[i]
		top++
	}

	return best
}
