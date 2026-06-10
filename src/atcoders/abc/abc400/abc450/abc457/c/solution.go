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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([][]int, n)
	for i := range n {
		var l int
		fmt.Fscan(reader, &l)
		a[i] = make([]int, l)
		for j := range l {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	c := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i])
	}
	return solve(k, a, c)
}

func solve(k int, a [][]int, c []int) int {
	for i, cur := range a {
		n := len(cur)
		if k <= n*c[i] {
			k = (k - 1) % n
			return cur[k]
		}
		k -= n * c[i]
	}
	return -1
}
