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
	var n, m, T int
	fmt.Fscan(reader, &n, &m, &T)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &b[i])
	}
	return solve(T, a, b)
}

const inf = 1 << 60

func solve(T int, a []int, b []int) int {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}

	m := len(b)
	for i := range m {
		b[i]--
	}

	calc := func(t int, i int) int {
		if T-t < 0 {
			return -inf
		}
		return (T - t) / a[i]
	}

	var ans int

	var lo int
	for j := range n {
		if lo+1 < m && b[lo+1] <= j {
			lo++
		}

		if b[lo] <= j {
			k := calc(sum[j]-sum[b[lo]], j) + j - b[lo]

			ans = max(ans, k)
			if lo+1 < m {
				// b[lo+1] > j
				k := calc(sum[b[lo+1]+1]-sum[j+1], j) + b[lo+1] - j
				ans = max(ans, k)
			}
		} else {
			k := calc(sum[b[lo]+1]-sum[j+1], j) + b[lo] - j
			ans = max(ans, k)
		}
	}

	return ans
}
