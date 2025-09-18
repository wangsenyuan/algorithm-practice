package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, x)
}

func solve(a []int, x int) []int {
	x--
	n := len(a)
	w := slices.Min(a)
	if a[x] == w {
		// 那么就是从x开始的
		sum := n * a[x]
		for i := range n {
			if i != x {
				a[i] -= a[x]
			}
		}
		a[x] = sum
		return a
	}
	var sum int
	y := x
	for a[y] != w {
		sum++
		y = (y + n - 1) % n
	}

	sum += n * a[y]
	for i := range n {
		if i != y {
			a[i] -= a[y]
		}
		if y < x && i > y && i <= x {
			a[i]--
		} else if y > x && (y < i || i <= x) {
			a[i]--
		}
	}
	a[y] = sum
	return a
}
