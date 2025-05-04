package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	res := solve(n, m)
	fmt.Println(res)
}

func solve(n int, m int) int {
	return gcd(n-1, m-1) + 1
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
