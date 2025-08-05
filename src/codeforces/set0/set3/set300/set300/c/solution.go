package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var a, b, n int
	fmt.Fscan(reader, &a, &b, &n)
	return solve(a, b, n)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
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

const N = 10_000_000 + 7

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(F[i-1], i)
	}
	I[N-1] = pow(F[N-1], mod-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(I[i+1], i+1)
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(a int, b int, n int) int {

	checkGood := func(num int) bool {
		for num > 0 {
			r := num % 10
			if r != a && r != b {
				return false
			}
			num /= 10
		}
		return true
	}

	var res int

	for i := 0; i <= n; i++ {
		// i个a组成
		sum := i*a + (n-i)*b
		if checkGood(sum) {
			res = add(res, nCr(n, i))
		}
	}
	return res
}
