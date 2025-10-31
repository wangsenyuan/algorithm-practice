package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var l, r int
	fmt.Fscan(reader, &l, &r)
	res := solve(l, r)
	fmt.Println(res)
}

func solve(l, r int) int {
	const inv2 = (mod + 1) / 2

	get := func(n int) int {
		return (f(n) + f((n+1)/2)) * inv2
	}

	ans := get(r) - get(l-1)
	return ((ans+mod)%mod + mod) % mod
}

const mod = 1_000_000_007

func kitamasa(coef, a []int, n int) (ans int) {
	if n < len(a) {
		return a[n]
	}
	k := len(coef)
	compose := func(a, b []int) []int {
		c := make([]int, k)
		for _, v := range a {
			for j, w := range b {
				c[j] = (c[j] + v*w) % mod
			}
			bk1 := b[k-1]
			for j := k - 1; j > 0; j-- {
				b[j] = (b[j-1] + bk1*coef[j]) % mod
			}
			b[0] = bk1 * coef[0] % mod
		}
		return c
	}

	resC := make([]int, k)
	resC[0] = 1
	c := make([]int, k)
	c[1] = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			resC = compose(c, resC)
		}
		c = compose(c, slices.Clone(c))
	}

	for i, c := range resC {
		ans = (ans + c*a[i]) % mod
	}
	return ans % mod
}

func f(n int) int {
	if n == 0 {
		return 0
	}
	coef := []int{-3, 3, 1}
	a := []int{4, 12, 26}
	return kitamasa(coef, a, n-1)
}
