package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	res := solve(n)
	fmt.Println(res)
}

func solve(n int) int {
	if n <= 2 {
		return 1
	}
	if n <= 4 {
		return 2
	}

	return (n-1)/2 + 1
}
