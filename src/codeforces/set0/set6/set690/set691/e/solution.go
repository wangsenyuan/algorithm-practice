package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

func solve(a []int, k int) int {
	n := len(a)
	tr := NewMat(n, n)
	for i := range n {
		for j := range n {
			if bits.OnesCount(uint(a[i]^a[j]))%3 == 0 {
				tr[i][j] = 1
			}
		}
	}
	res := NewMat(n, 1)
	for i := range n {
		res[i][0] = 1
	}
	res = matPow(tr, k-1, res)
	var ans int
	for i := range n {
		ans = add(ans, res[i][0])
	}
	return ans
}

const mod = 1e9 + 7

func add(a, b int) int {
	return (a + b) % mod
}

func mul(a, b int) int {
	return (a * b) % mod
}

type mat [][]int

func NewMat(n int, m int) mat {
	res := make(mat, n)
	for i := range n {
		res[i] = make([]int, m)
	}
	return res
}

func (this mat) mul(that mat) mat {
	n := len(this)
	m := len(this[0])
	k := len(that[0])
	res := NewMat(n, k)
	for j := range m {
		for i := range n {
			if this[i][j] != 0 {
				for l := range k {
					res[i][l] = add(res[i][l], mul(this[i][j], that[j][l]))
				}
			}
		}
	}
	return res
}

func matPow(a mat, n int, res mat) mat {
	for n > 0 {
		if n&1 == 1 {
			res = a.mul(res)
		}
		a = a.mul(a)
		n >>= 1
	}
	return res
}
