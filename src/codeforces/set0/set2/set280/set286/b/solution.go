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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func solve(n int) []int {
	a := make([]int, 3*n+5)
	for i := range n + 1 {
		a[i] = i
	}

	for i := 2; i <= n; i++ {
		for j := 2*i - 1; a[j] > 0; j += i {
			a[i-1], a[j] = a[j], a[i-1]
		}
		a[i-1], a[i-1+n] = a[i-1+n], a[i-1]
	}

	return a[n : 2*n]
}
