package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) (a []int, res string) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res = solve(a)
	return
}

func solve(a []int) string {
	n := len(a)
	diff := make([]int, n)
	next := a[n-1]
	diff[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		diff[i] = 1
		if a[i] <= next {
			// next <= 2 * a[i]成立
			diff[i] *= -1
			diff[i+1] *= -1
			next = -a[i] + next
		} else {
			// next < a[i]
			diff[i+1] *= -1
			next = a[i] - next
		}
	}

	buf := make([]byte, n)
	for i := range n {
		if i > 0 {
			diff[i] *= diff[i-1]
		}
		if diff[i] == 1 {
			buf[i] = '+'
		} else {
			buf[i] = '-'
		}
	}
	return string(buf)
}
