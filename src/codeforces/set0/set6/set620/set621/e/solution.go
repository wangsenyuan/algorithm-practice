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
	var n, b, k, x int
	fmt.Fscan(reader, &n, &b, &k, &x)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(b, k, x, a)
}

func solve(b int, k int, x int, a []int) int {
	freq := make([]int, 10)
	for _, v := range a {
		freq[v]++
	}

	dp := NewMat(x, x)

	for u := range x {
		for i := 1; i <= 9; i++ {
			v := (u*10 + i) % x
			dp[v][u] = add(dp[v][u], freq[i])
		}
	}

	f0 := NewMat(x, 1)
	for i := range 10 {
		f0[i%x][0] += freq[i]
	}
	res := powMat(dp, b-1, f0)

	return res[k][0]
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
