package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func readMatrix(reader *bufio.Reader, n, m int) [][]int {
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return a
}

func drive(reader *bufio.Reader) bool {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := readMatrix(reader, n, m)
	b := readMatrix(reader, n, m)
	return solve(a, b)
}

func solve(a [][]int, b [][]int) bool {
	n := len(a)
	m := len(a[0])
	c := make([][]int, n)
	for i := range n {
		c[i] = make([]int, m)
		for j := range m {
			c[i][j] = a[i][j] ^ b[i][j]
		}
	}

	for i := range n {
		for j := range m {
			if c[i][j] == 1 {
				if j == m-1 || i == n-1 {
					return false
				}
				c[i][j] ^= 1
				c[i+1][j] ^= 1
				c[i][j+1] ^= 1
				c[i+1][j+1] ^= 1
			}
		}
	}

	return true
}
