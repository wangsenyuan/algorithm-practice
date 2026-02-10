package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var n, f int
		fmt.Fscan(reader, &n, &f)
		res := solve(n, f)
		fmt.Fprintln(writer, res)
	}
}

const mod = 1000000007

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

const N = 100005

var F [N]int

var I [N]int
var fs [N][]int
var mu [N]int

func init() {
	F[0] = 1
	for i := 1; i < N; i++ {
		F[i] = mul(i, F[i-1])
	}
	I[N-1] = inverse(F[N-1])
	for i := N - 2; i >= 0; i-- {
		I[i] = mul(i+1, I[i+1])
	}

	for i := 1; i < N; i++ {
		for j := i; j < N; j += i {
			fs[j] = append(fs[j], i)
		}
	}

	sq := make([]bool, N)
	for i := 2; i*i < N; i++ {
		t := i * i
		for j := t; j < N; j += t {
			sq[j] = true
		}
	}

	prime := make([]bool, N)
	for i := range N {
		prime[i] = true
	}
	prime[0] = false
	prime[1] = false
	for i := 2; i < N; i++ {
		if prime[i] {
			for j := i * 2; j < N; j += i {
				prime[j] = false
			}
		}
	}

	for i := 1; i < N; i++ {
		if sq[i] {
			continue
		}
		var parity int
		for _, j := range fs[i] {
			if prime[j] {
				parity++
			}
		}
		if parity&1 == 1 {
			mu[i] = -1
		} else {
			mu[i] = 1
		}
	}
}

func nCr(n int, r int) int {
	if n < r || r < 0 {
		return 0
	}
	return mul(F[n], mul(I[r], I[n-r]))
}

func solve(n, f int) int {
	var res int
	for _, x := range fs[n] {
		if x*f > n {
			break
		}
		if mu[x] == 0 {
			continue
		}
		// x * f <= n
		m := n / x

		cur := nCr(m-1, f-1)

		if mu[x] == 1 {
			res = add(res, cur)
		} else {
			res = sub(res, cur)
		}
	}

	return res
}
