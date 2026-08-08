package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
}

func solve(n int) int {
	if checkPrime(n) {
		return 1
	}
	if n%2 == 0 {
		return 2
	}

	if checkPrime(n - 2) {
		return 2
	}

	return 3
}

func checkPrime(n int) bool {
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
