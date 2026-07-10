package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 998244353

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, ans := range drive(reader) {
		fmt.Fprintln(writer, ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var tc int
	fmt.Fscan(reader, &tc)
	res := make([]int, tc)
	for i := range tc {
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &a[j])
		}
		res[i] = solve(a)
	}
	return res
}

func solve(a []int) int {
	n := len(a)
	h := (n + 1) / 2

	fac, ifac := make([]int, n+1), make([]int, n+1)
	fac[0] = 1
	for i := 1; i <= n; i++ {
		fac[i] = mul(fac[i-1], i)
	}
	ifac[n] = pow(fac[n], mod-2)
	for i := n; i > 0; i-- {
		ifac[i-1] = mul(ifac[i], i)
	}

	nCr := func(n int, r int) int {
		if r < 0 || r > n {
			return 0
		}
		return mul(fac[n], mul(ifac[r], ifac[n-r]))
	}

	var ans = 1
	var available int
	for row := n; row > 0; row-- {
		if row <= h {
			available += 2
			if n%2 == 1 && row == h {
				available--
			}
		}
		ans = mul(ans, nCr(available, a[row-1]))
		available -= a[row-1]
		if available < 0 {
			return 0
		}
	}
	if available != 0 {
		return 0
	}
	return ans
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}
