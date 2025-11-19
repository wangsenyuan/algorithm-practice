package main

import "fmt"

func main() {
	var C, W, H int
	fmt.Scan(&C, &W, &H)
	fmt.Println(solve(C, W, H))
}

func solve(C int, W int, H int) int {
	if C <= W {
		return pow(H+1, C)
	}
	// C > W
	mat := NewMat(W+1, W+1)
	for i := range W + 1 {
		mat[0][i] = 1
		if i > 0 {
			mat[i][i-1] = H
		}
	}
	f := NewMat(W+1, 1)
	f[0][0] = 1

	res := powMat(mat, C, f)
	var ans int
	for i := range W + 1 {
		ans = add(ans, res[i][0])
	}
	return ans
}

const mod = 1e6 + 3

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return (a * b) % mod
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
	for i := range res {
		res[i] = make([]int, m)
	}
	return res
}

func (a mat) mul(b mat) mat {
	n := len(a)
	m := len(a[0])
	k := len(b[0])
	res := NewMat(n, k)
	for l := range m {
		for i := range n {
			if a[i][l] == 0 {
				continue
			}
			for j := range k {
				res[i][j] = add(res[i][j], mul(a[i][l], b[l][j]))
			}
		}
	}
	return res
}

func powMat(a mat, n int, res mat) mat {
	for n > 0 {
		if n&1 == 1 {
			res = a.mul(res)
		}
		a = a.mul(a)
		n >>= 1
	}
	return res
}
