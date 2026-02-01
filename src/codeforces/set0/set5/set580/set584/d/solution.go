package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int

	fmt.Fscan(reader, &n)
	res := solve(n)

	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int) []int {
	if checkPrime(n) {
		return []int{n}
	}

	if checkPrime(n - 2) {
		return []int{2, n - 2}
	}

	// k is even
	k := n - 3

	if checkPrime(k / 2) {
		return []int{3, k / 2, k / 2}
	}

	i, j := k/2-1, k/2+1
	for !checkPrime(i) || !checkPrime(j) {
		i--
		j++
	}
	return []int{3, i, j}
}

func checkPrime(n int) bool {
	if n == 2 {
		return true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
