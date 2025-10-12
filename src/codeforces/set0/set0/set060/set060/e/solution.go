package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, x, y, p int
	fmt.Fscan(reader, &n, &x, &y, &p)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(reader, &a[i])
	}
	return solve(x, y, a, p)
}

func solve(x int, y int, a []int, p int) int {
	n := len(a)
	if n == 1 {
		return a[0] % p
	}
	var sum int
	for _, v := range a {
		sum = add(sum, v, p)
	}

	f := mat{{1, 1}, {1, 0}}
	r := mat{{a[n-1] % p}, {a[n-2] % p}}

	f = matPow(f, x, r, p)
	big := f[0][0]

	calc := func(s int, c int, n int) int {
		m := mat{{3, 1}, {0, 1}}
		r := mat{{s}, {sub(0, c, p)}}
		m = matPow(m, n, r, p)
		return m[0][0]
	}

	s1 := calc(sum, a[0]+a[n-1], x)

	s2 := calc(s1, a[0]+big, y)

	return s2
}

type mat [][]int

func add(a, b, mod int) int {
	return (a + b) % mod
}
func sub(a, b, mod int) int {
	b %= mod
	return add(a, mod-b, mod)
}

func mul(a, b, mod int) int {
	return a * b % mod
}

func NewMat(n int, m int) mat {
	res := make(mat, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	return res
}

func (a mat) mul(b mat, mod int) mat {
	n := len(a)
	m := len(a[0])
	k := len(b[0])
	res := NewMat(n, k)
	for j := range m {
		for i := range n {
			if a[i][j] != 0 {
				for l := range k {
					res[i][l] = add(res[i][l], mul(a[i][j], b[j][l], mod), mod)
				}
			}
		}
	}
	return res
}

func matPow(a mat, n int, r0 mat, mod int) mat {
	res := r0
	for n > 0 {
		if n&1 == 1 {
			res = a.mul(res, mod)
		}
		a = a.mul(a, mod)
		n >>= 1
	}
	return res
}
