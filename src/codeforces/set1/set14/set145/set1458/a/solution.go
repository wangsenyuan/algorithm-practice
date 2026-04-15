package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	res := solve(a, b)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func solve(a []int, b []int) []int {
	var g int
	n := len(a)
	for i := range n - 1 {
		g = gcd(g, abs(a[i+1]-a[i]))
	}
	m := len(b)
	res := make([]int, m)
	for i, v := range b {
		res[i] = gcd(g, v+a[0])
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func abs(num int) int {
	return max(num, -num)
}
