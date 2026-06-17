package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n+1)
	for i := 0; i <= n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

const mod = 998244353

func mul(nums ...int) int {
	res := 1
	for _, num := range nums {
		res *= num
		res %= mod
	}
	return res
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

const N = 55

var F [N]int
var I [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[N-1] = pow(F[N-1], mod-2)
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}
}

func nCr(n int, r int) int {
	return mul(F[n], I[r], I[n-r])
}

func solve(a []int) int {
	n := len(a)
	var sum int
	for _, v := range a {
		sum += v
	}
	k := sum / (n - 1)
	// 至少要经历k轮

	for i := 1; i < n; i++ {
		if a[i] >= k {
			a[i] -= k
		} else {
			a[0] -= (k - a[i])
			a[i] = 0
		}
	}
	if a[0] < 0 {
		return 0
	}
	// b[i]表示经过k轮后剩余的, 然后最多还剩一轮,
	if slices.Max(a[1:]) > 1 {
		return 0
	}
	var z int
	for i := 1; i < n; i++ {
		if a[i] == 0 {
			z++
		}
	}
	if a[0] > z {
		return 0
	}
	// b[0] <= z
	x := z - a[0]
	// 选出x个人在末尾, 它们可以之间随便排列, 其他的n-x个人, 也可以随便排列
	return mul(nCr(z, x), F[x], F[n-1-x])
}
