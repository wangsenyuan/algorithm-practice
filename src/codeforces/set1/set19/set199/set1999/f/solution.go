package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

const N = 200_000 + 10

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[N-1] = inverse(F[N-1])
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

const MOD = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= MOD {
		a -= MOD
	}
	return a
}
func mul(a, b int) int {
	return a * b % MOD
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
	return pow(a, MOD-2)
}

func nCr(n int, r int) int {
	if n < 0 || n < r {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(a []int, k int) int {
	sort.Ints(a)

	var res int

	h := k / 2

	n := len(a)
	for i := h; i+h < n; i++ {
		if a[i] == 1 {
			res = add(res, mul(nCr(i, h), nCr(n-i-1, h)))
		}
	}

	return res
}
