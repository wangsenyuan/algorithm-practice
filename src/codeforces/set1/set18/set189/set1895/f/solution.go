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
		var n, x, k int
		fmt.Fscan(reader, &n, &x, &k)
		res := solve(n, x, k)
		fmt.Fprintln(writer, res)
	}
}

func solve(n int, x int, k int) int {
	// (x + k) * (2k + 1)^{n-1}
	ans := mul(add(x, k), pow(add(mul(2, k), 1), n-1))
	if x > 0 {
		m := NewMat(x, x)
		for j := range m {
			for v := max(j-k, 0); v <= min(j+k, x-1); v++ {
				m[j][v] = 1
			}
		}
		f1 := NewMat(x, 1)
		for j := range f1 {
			f1[j][0] = 1
		}
		fn := powMat(m, n-1, f1)
		for _, row := range fn {
			ans = sub(ans, row[0])
		}
	}

	return ans
}

const mod = 1e9 + 7

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

type mat [][]int

func NewMat(n int, m int) mat {
	res := make(mat, n)
	for i := range n {
		res[i] = make([]int, m)
	}
	return res
}

func (a mat) mul(b mat) mat {
	n := len(a)
	m := len(a[0])
	k := len(b[0])
	res := NewMat(n, k)
	for j := range m {
		for i := range n {
			for u := range k {
				res[i][u] = add(res[i][u], mul(a[i][j], b[j][u]))
			}
		}
	}
	return res
}

func identityMat(n int) mat {
	res := NewMat(n, n)
	for i := range n {
		res[i][i] = 1
	}
	return res
}

func powMat(a mat, b int, res mat) mat {
	for b > 0 {
		if b&1 == 1 {
			res = a.mul(res)
		}
		a = a.mul(a)
		b >>= 1
	}
	return res
}
