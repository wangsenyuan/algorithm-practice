package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}	

func drive(reader *bufio.Reader) int {
	var n, d int
	fmt.Fscan(reader, &n, &d)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(d, a)
}

const mod = 998244353

func mul(a, b int) int {
	return a * b % mod
}

func pow(a int, b int) int {
	r := 1
	for b > 0 {
		if b&1 == 1 {
			r = mul(r, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return r
}

func div(a, b int) int {
	return mul(a, pow(b, mod-2))
}

func solve(d int, a []int) int {
	slices.Sort(a)

	n := len(a)
	f := make([]int, n+1)
	f[0] = 1
	for i := range n {
		f[i+1] = mul(i+1, f[i])
	}

	res := 1
	var l int
	var cnt int
	for r, v := range a {
		if r > 0 && v != a[r-1] {
			res = div(res, f[cnt])
			cnt = 0
		}
		cnt++

		for l < r && a[l] < v-d {
			l++
		}
		// a[l] >= v - d, 所以, v可以放置到l的后面去
		res = mul(res, r-l+1)
	}
	res = div(res, f[cnt])
	return res
}
