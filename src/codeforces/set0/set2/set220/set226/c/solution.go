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
	var m, l, r, k int
	fmt.Fscan(reader, &m, &l, &r, &k)
	return solve(m, l, r, k)
}

func solve(m int, l int, r int, k int) int {
	if m == 1 {
		return 0
	}

	add := func(a, b int) int {
		return (a + b) % m
	}

	mul := func(a, b int) int {
		return a * b % m
	}

	type mat [][]int

	newMat := func(n, k int) mat {
		res := make(mat, n)
		for i := range n {
			res[i] = make([]int, k)
		}
		return res
	}

	mulMat := func(a, b mat) mat {
		n := len(a)
		m := len(a[0])
		k := len(b[0])
		res := newMat(n, k)
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

	powMat := func(a mat, n int, res mat) mat {
		for n > 0 {
			if n&1 == 1 {
				res = mulMat(a, res)
			}
			a = mulMat(a, a)
			n >>= 1
		}
		return res
	}

	l--
	w := -1
	for i := 1; i*i <= r; i++ {
		if j := r / i; r/j-l/j >= k {
			w = j
			break
		}
		if r/i-l/i >= k {
			w = i
		}
	}

	a := newMat(2, 2)
	a[0][0] = 1
	a[0][1] = 1
	a[1][0] = 1
	res := newMat(2, 1)
	res[0][0] = 1
	res = powMat(a, w-1, res)

	return res[0][0] % m
}
