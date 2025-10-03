package main

import (
	"fmt"
	"os"
)

func main() {
	var n, m int
	fmt.Fscan(os.Stdin, &n, &m)
	res := solve(n, m)
	fmt.Println(res)
}

const mod = 1000000007

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
	for w := range m {
		for i := range n {
			if a[i][w] != 0 {
				for j := range k {
					res[i][j] = add(res[i][j], mul(a[i][w], b[w][j]))
				}
			}
		}
	}
	return res
}

func matPow(a mat, n int, r0 mat) mat {
	res := r0
	for n > 0 {
		if n&1 == 1 {
			res = a.mul(res)
		}
		a = a.mul(a)
		n >>= 1
	}
	return res
}

func solve(n int, m int) int {
	a := NewMat(m, m)
	// f(i) = f(i-1) + f(i - m)
	a[0][0] = 1
	a[0][m-1] = 1
	for i := 0; i < m-1; i++ {
		a[i+1][i] = 1
	}
	b := NewMat(m, 1)
	b[0][0] = 1
	res := matPow(a, n, b)
	return res[0][0]
}
