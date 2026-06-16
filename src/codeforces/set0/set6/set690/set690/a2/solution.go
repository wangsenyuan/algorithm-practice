package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	if n&1 == 1 {
		return (n - 1) / 2
	}

	// n is even
	var h int
	for 1<<(h+1) <= n {
		h++
	}

	return (n - (1 << h)) / 2
}
