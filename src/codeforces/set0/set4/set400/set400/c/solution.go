package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n, m, x, y, z, p int
	fmt.Fscan(reader, &n, &m, &x, &y, &z, &p)
	candies := make([][]int, p)
	for i := range p {
		candies[i] = make([]int, 2)
		fmt.Fscan(reader, &candies[i][0], &candies[i][1])
	}
	return solve(n, m, x, y, z, candies)
}

func solve(n int, m int, x int, y int, z int, candies [][]int) [][]int {
	x %= 4
	y %= 2
	z %= 4

	rotate := func() {
		for i := range candies {
			r, c := candies[i][0], candies[i][1]
			candies[i][0] = c
			candies[i][1] = n - r + 1
		}
	}

	for range x {
		rotate()
		n, m = m, n
	}

	if y == 1 {
		for i, cur := range candies {
			c := cur[1]
			candies[i][1] = m - c + 1
		}
	}

	if z != 0 {
		z = 4 - z
		for range z {
			rotate()
			n, m = m, n
		}
	}

	return candies
}
