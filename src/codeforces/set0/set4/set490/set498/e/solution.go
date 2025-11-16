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
	w := make([]int, 7)
	for i := range 7 {
		fmt.Fscan(reader, &w[i])
	}
	return solve(w)
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

func solve(w []int) int {

	calc := func(h int) [][]int {
		// dp[s1][s2]
		S := 1 << h
		dp := make([][]int, S)
		for i := range S {
			dp[i] = make([]int, S)
		}

		if h == 1 {
			dp[0][0] = 1
			dp[1][0] = 1
			dp[0][1] = 1
			return dp
		}

		for s1 := range S {
			for s2 := range S {
				for x := range 1 << (h - 1) {
					ok := true
					for i := range h - 1 {
						w := (x >> i) & 1
						// 这里要处理的是里面横线的状态
						t := 1
						if i > 0 {
							t = (x >> (i - 1)) & 1
						}
						l := (s1 >> i) & 1
						r := (s2 >> i) & 1
						if t+l+r+w == 4 {
							ok = false
							break
						}
						b := 1
						if i+1 < h-1 {
							b = (x >> (i + 1)) & 1
						}
						l = (s1 >> (i + 1)) & 1
						r = (s2 >> (i + 1)) & 1
						if w+b+l+r == 4 {
							ok = false
							break
						}
					}

					if ok {
						dp[s1][s2]++
					}
				}
			}
		}
		return dp
	}

	type pair struct {
		first  int
		second int
	}

	var arr []pair
	for i := range 7 {
		if w[i] != 0 {
			arr = append(arr, pair{first: i + 1, second: w[i]})
		}
	}

	var f0 mat

	for i, cur := range arr {
		h, w := cur.first, cur.second
		dp := calc(h)

		f1 := NewMat(1<<h, 1)
		if i > 0 {
			h0 := arr[i-1].first
			diff := h - h0
			pad := 1<<diff - 1
			for s1 := range 1 << h0 {
				s2 := s1<<diff + pad
				f1[s2][0] = f0[s1][0]
			}
		} else {
			f1[1<<h-1][0] = 1
		}

		f0 = powMat(dp, w, f1)
	}

	h1 := arr[len(arr)-1].first
	return f0[1<<h1-1][0]
}
