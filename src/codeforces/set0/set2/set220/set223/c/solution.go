package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, k, a)
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
	return int(int64(a) * int64(b) % mod)
}

func pow(a, e int) int {
	r := 1
	for e > 0 {
		if e&1 != 0 {
			r = mul(r, a)
		}
		a = mul(a, a)
		e >>= 1
	}
	return r
}

func inv(a int) int {
	return pow(a, mod-2)
}

func solve(n, k int, a []int) []int {
	// After k prefix-sum ops, a[j] contributes C(k-1+d, d) to position j+d.
	// C[d] = C(k-1+d, d); C[0]=1, C[d+1]=C[d]*(k+d)/(d+1).
	C := make([]int, n)
	C[0] = 1
	for d := 0; d+1 < n; d++ {
		C[d+1] = mul(C[d], mul(add(k%mod, d), inv(d+1)))
	}

	ans := make([]int, n)
	for i := range n {
		var s int
		for j := 0; j <= i; j++ {
			s = add(s, mul(a[j]%mod, C[i-j]))
		}
		ans[i] = s
	}
	return ans
}
