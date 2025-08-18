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
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) int {
	n := len(a)
	stack := make([]int, n)
	var top int
	prev := make([]int, n)
	for i := range n {
		prev[i] = -1
		for top > 0 && a[stack[top-1]] < a[i] {
			top--
		}
		if top > 0 {
			prev[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}
	top = 0
	next := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		next[i] = n
		for top > 0 && a[stack[top-1]] <= a[i] {
			top--
		}
		if top > 0 {
			next[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}

	pos := make([]int, 30)
	for i := range 30 {
		pos[i] = -1
	}
	L := make([]int, n)
	for i := range n {
		lb := prev[i]
		for j := range 30 {
			if (a[i]>>j)&1 == 0 {
				lb = max(lb, pos[j])
			} else {
				pos[j] = i
			}
		}
		L[i] = lb
	}
	for i := range 30 {
		pos[i] = n
	}
	R := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		rb := next[i]
		for j := range 30 {
			if (a[i]>>j)&1 == 0 {
				rb = min(rb, pos[j])
			} else {
				pos[j] = i
			}
		}
		R[i] = rb
	}

	ans := n * (n - 1) / 2

	for i := range n {
		x := i - L[i]
		y := R[i] - i
		ans -= x*y - 1
	}

	return ans
}
