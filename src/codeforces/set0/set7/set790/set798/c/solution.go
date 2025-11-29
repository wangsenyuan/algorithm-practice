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
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	res := solve(a)
	fmt.Println("YES")
	fmt.Println(res)
}

func solve(a []int) int {
	var g int
	for _, v := range a {
		g = gcd(g, v)
	}
	if g > 1 {
		return 0
	}
	n := len(a)

	var res int
	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			continue
		}
		res++
		if i == n-1 || a[i+1]%2 == 0 {
			res++
		}
		i++
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
