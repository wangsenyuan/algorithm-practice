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

func pow(a, b int) int {
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

func inverse(a int) int {
	return pow(a, mod-2)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	return solve(n, m, a)
}

func solve(n int, m int, a [][]int) int {
	cnt := make([]int, n+1)

	var sum int

	for j := range m {
		clear(cnt)

		for i := range n {
			cnt[n+1-a[i][j]]++
		}

		cur := 1
		var rem int
		for i := range n {
			rem += cnt[i]
			if rem == 0 {
				cur = 0
				break
			}
			cur = mul(cur, rem)
			rem = max(0, rem-1)
		}
		sum = add(sum, cur)
	}

	pn := 1
	for i := range n {
		pn = mul(pn, i+1)
	}
	res := sub(mul(pn, m), sum)
	res = mul(res, inverse(pn))
	return res
}
