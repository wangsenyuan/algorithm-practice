package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x, y, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &x, &y, &m)
	fmt.Println(solve(x, y, m))
}

func solve(x int, y int, m int) int {
	if max(x, y) >= m {
		return 0
	}
	// max(x, y) < m
	if max(x, y) <= 0 {
		return -1
	}
	x, y = max(x, y), min(x, y)
	// 要想将y变成 > 0的数
	var res int
	if y < 0 {
		// n * x + y > 0
		// n * x > -y
		res = (-y + x - 1) / x
		y += res * x
	}

	// x > 0 and y > 0
	// 每次都用 x + y 去替换小的数
	for x < m {
		y = x + y
		x, y = y, x
		res++
	}

	return res
}
