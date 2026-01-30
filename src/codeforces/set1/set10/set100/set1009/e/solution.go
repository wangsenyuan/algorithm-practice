package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const mod = 998244353

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(a []int) int {

	n := len(a)
	pw := make([]int, n+1)
	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = add(pw[i-1], pw[i-1])
	}

	cur := mul(a[0], pw[n-1])
	var ans int
	for i := range n {
		ans = add(ans, cur)
		if i < n-1 {
			cur = sub(cur, mul(a[i], pw[n-2-i]))
			cur = add(cur, mul(a[i+1], pw[n-2-i]))
		}
	}

	return ans
}
