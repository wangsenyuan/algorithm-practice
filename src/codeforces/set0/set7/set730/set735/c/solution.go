package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
}

func solve(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	// 1, 2
	a, b := 2, 3

	for i := 3; ; i++ {
		c := a + b
		if c > n {
			return i - 1
		}
		a, b = b, c
	}
}
