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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(k, p)
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

func inv(a int) int {
	return pow(a, mod-2)
}

func solve(k int, p []int) int {
	n := len(p)

	pref := make(BIT, n+3)

	var sum int
	for i := range n {
		sum = add(sum, pref.rangeQuery(p[i], n))
		pref.update(p[i], 1)
	}

	clear(pref)

	var win int

	sum = add(sum, mul(k*(k-1)/2%mod, inv(2)))

	for i := range k {
		win = add(win, pref.rangeQuery(p[i], n))
		pref.update(p[i], 1)
	}

	var res int

	for i := k; i <= n; i++ {
		res = add(res, sub(sum, win))
		if i == n {
			break
		}
		pref.update(p[i-k], -1)
		win = sub(win, pref.rangeQuery(0, p[i-k]))
		win = add(win, pref.rangeQuery(p[i], n))
		pref.update(p[i], 1)
	}

	res = mul(res, inv(n-k+1))

	return res
}

type BIT []int

func (bit BIT) update(i int, v int) {
	i++
	for i < len(bit) {
		bit[i] += v
		i += i & -i
	}
}

func (bit BIT) query(i int) int {
	i++
	var res int
	for i > 0 {
		res += bit[i]
		i -= i & -i
	}
	return res
}

func (bit BIT) rangeQuery(l int, r int) int {
	if l > r {
		return 0
	}
	return bit.query(r) - bit.query(l)
}
