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
	var n, k int
	fmt.Fscan(reader, &n, &k)
	segments := make([][]int, n)
	for i := range n {
		var a, b, c int
		fmt.Fscan(reader, &a, &b, &c)
		segments[i] = []int{a, b, c}
	}
	return solve(k, segments)
}

func solve(k int, segments [][]int) int {
	f := NewMat(16, 1)
	f[0][0] = 1
	for w, cur := range segments {
		l, r, y := cur[0], cur[1], cur[2]
		mat := NewMat(16, 16)
		for i := range y + 1 {
			for j := max(0, i-1); j <= min(y, i+1); j++ {
				mat[i][j] = 1
			}
		}

		if w == len(segments)-1 {
			r = k
		}

		f = powMat(mat, r-l, f)
	}
	return f[0][0]
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
	for l := range m {
		for i := range n {
			for j := range k {
				if a[i][l] != 0 && b[l][j] != 0 {
					res[i][j] = add(res[i][j], mul(a[i][l], b[l][j]))
				}
			}
		}
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
