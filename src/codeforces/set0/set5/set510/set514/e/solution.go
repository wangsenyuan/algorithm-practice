package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	d := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &d[i])
	}
	return solve(x, d)
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

func solve(x int, d []int) int {
	dp := make([]int, 101)
	dp[0] = 1

	cnt := make([]int, 101)
	for _, v := range d {
		cnt[v]++
	}

	for i := 1; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = add(dp[i], mul(cnt[j], dp[i-j]))
		}
	}

	if x <= 100 {
		var res int
		for i := range x + 1 {
			res = add(res, dp[i])
		}
		return res
	}

	B := make(mat, 101)
	for i := range 101 {
		B[i] = make([]int, 101)
		B[i][99] = cnt[100-i]
		B[i][100] = cnt[100-i]
		if i > 0 {
			B[i][i-1] = 1
		}
	}

	B[100][99] = 0
	B[100][100] = 1

	B = B.pow(x - 100)

	A := make(mat, 1)
	A[0] = make([]int, 101)
	sum := dp[0]
	for i := 0; i < 100; i++ {
		A[0][i] = dp[i+1]
		sum = add(sum, dp[i+1])
	}
	A[0][100] = sum

	res := A.mul(B)
	return res[0][100]
}

type mat [][]int

func (this mat) mul(that mat) mat {
	n := len(this)
	m := len(this[0])
	// m = len(that)
	k := len(that[0])
	res := make(mat, n)
	for i := range res {
		res[i] = make([]int, k)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for l := 0; l < m; l++ {
				res[i][j] = add(res[i][j], mul(this[i][l], that[l][j]))
			}
		}
	}
	return res
}

func identity(n int) mat {
	res := make(mat, n)
	for i := range res {
		res[i] = make([]int, n)
		res[i][i] = 1
	}
	return res
}

func (this mat) pow(n int) mat {
	// len(mat) = len(mat[0])
	if n == 0 {
		return identity(len(this))
	}
	if n == 1 {
		return this
	}
	half := this.pow(n / 2)
	res := half.mul(half)
	if n&1 == 1 {
		res = res.mul(this)
	}
	return res
}
