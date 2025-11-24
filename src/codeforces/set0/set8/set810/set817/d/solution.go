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
	L := make([]int, n)

	stack := make([]int, n)
	var top int
	for i, x := range a {
		for top > 0 && a[stack[top-1]] < x {
			top--
		}
		if top == 0 {
			L[i] = -1
		} else {
			L[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}

	var res int
	top = 0
	for i := n - 1; i >= 0; i-- {
		for top > 0 && a[stack[top-1]] <= a[i] {
			top--
		}
		R := n
		if top > 0 {
			R = stack[top-1]
		}
		w := (R - i) * (i - L[i])
		res += a[i] * w
		stack[top] = i
		top++
	}
	top = 0
	for i, x := range a {
		for top > 0 && a[stack[top-1]] > x {
			top--
		}
		L[i] = -1
		if top > 0 {
			L[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}

	top = 0
	for i := n - 1; i >= 0; i-- {
		for top > 0 && a[stack[top-1]] >= a[i] {
			top--
		}
		R := n
		if top > 0 {
			R = stack[top-1]
		}
		w := (R - i) * (i - L[i])
		res -= a[i] * w
		stack[top] = i
		top++
	}

	return res
}
