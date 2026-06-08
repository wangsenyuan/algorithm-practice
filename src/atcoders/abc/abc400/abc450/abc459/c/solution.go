package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	ops := make([][]int, m)
	for i := range m {
		ops[i] = make([]int, 2)
		fmt.Fscan(reader, &ops[i][0], &ops[i][1])
	}
	return solve(n, ops)
}

func solve(n int, ops [][]int) []int {
	m := len(ops)

	c := make([]int, m+1)
	a := make([]int, n+1)
	var mn int
	var res []int
	for _, cur := range ops {
		if cur[0] == 1 {
			x := cur[1]
			a[x]++
			c[a[x]]++
			if c[a[x]] == n {
				mn = a[x]
			}
		} else {
			y := cur[1]
			if y+mn > m {
				res = append(res, 0)
			} else {
				res = append(res, c[y+mn])
			}
		}
	}

	return res
}
