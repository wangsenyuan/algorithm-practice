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
	var n, l, m int
	fmt.Fscan(reader, &n, &l, &m)
	first := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &first[i])
	}
	second := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &second[i])
	}
	third := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &third[i])
	}
	return solve(l, m, first, second, third)
}

func solve(L int, m int, first []int, second []int, third []int) int {
	//n := len(first)

	getFreq := func(arr []int) []int {
		res := make([]int, m)
		for _, v := range arr {
			res[v%m]++
		}
		return res
	}
	f1 := getFreq(first)
	f2 := getFreq(second)
	f3 := getFreq(third)
	tf := make(mat, m)
	for i := range m {
		tf[i] = make([]int, m)
	}

	for x := range m {
		for y := range m {
			tf[x][y] = add(tf[x][y], f2[(y-x+m)%m])
		}
	}
	tf = mat_pow(tf, L-2)
	dp := make([]int, m)
	for x := range m {
		for y := range m {
			dp[(x+y)%m] = add(dp[(x+y)%m], mul(f1[x], tf[x][y]))
		}
	}
	var res int
	for x := range m {
		for y := range m {
			if (x+y)%m == 0 {
				res = add(res, mul(dp[x], f3[y]))
			}
		}
	}
	return res
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

func (this mat) mul(that mat) mat {
	h := len(this)
	w := len(this[0])
	k := len(that[0])
	res := make(mat, h)
	for i := range h {
		res[i] = make([]int, k)
	}
	for l := range w {
		for i := range h {
			for j := range k {
				res[i][j] = add(res[i][j], mul(this[i][l], that[l][j]))
			}
		}
	}
	return res
}

func identity(n int) mat {
	res := make(mat, n)
	for i := range n {
		res[i] = make([]int, n)
		res[i][i] = 1
	}
	return res
}

func mat_pow(a mat, n int) mat {
	if n == 0 {
		return identity(len(a))
	}
	if n == 1 {
		return a
	}
	res := mat_pow(a, n/2)
	if n%2 == 1 {
		res = res.mul(a)
	}
	return res
}
