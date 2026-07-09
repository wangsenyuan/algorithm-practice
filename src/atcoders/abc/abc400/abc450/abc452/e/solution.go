package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
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
	return solve(a, b)
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
	if a < b {
		a += mod
	}
	return a - b
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(a, b []int) int {
	a = append([]int{0}, a...)
	b = append([]int{0}, b...)
	n := len(a)
	m := len(b)

	// 然后计算 j > i 的部分, 这些就是 a[i] * i * sum(b[j])
	sum1 := make([]int, n+1)
	sum2 := make([]int, n+1)
	for i := range n {
		sum1[i+1] = add(sum1[i], a[i])
		sum2[i+1] = add(sum2[i], mul(a[i], i))
	}

	var ans int

	for j := 1; j < m; j++ {
		var cur int
		for i := 0; i*j < n; i++ {
			l := i * j
			r := min(n, l+j)
			cur = add(cur, sub(sum2[r], sum2[l]))
			cur = sub(cur, mul(sub(sum1[r], sum1[l]), mul(i, j)))
		}

		ans = add(ans, mul(cur, b[j]))
	}

	return ans
}
