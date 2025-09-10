package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	obstacles := make([][]int, n)
	for i := 0; i < n; i++ {
		obstacles[i] = make([]int, 3)
		fmt.Fscan(reader, &obstacles[i][0], &obstacles[i][1], &obstacles[i][2])
	}
	return solve(m, obstacles)
}

func solve(m int, obstacles [][]int) int {

	// 这里要把这些障碍物重新组织一下？
	// row, l, r, 变成从左到右一段一段的
	var all []int
	for _, cur := range obstacles {
		all = append(all, cur[1]-1, cur[2])
	}
	all = append(all, 1)
	all = append(all, m)
	slices.Sort(all)
	all = slices.Compact(all)

	n := len(all)
	diff := make([][3]int, n+1)
	for _, cur := range obstacles {
		i := cur[0] - 1
		l := sort.SearchInts(all, cur[1]-1) + 1
		r := sort.SearchInts(all, cur[2])
		diff[l][i]++
		diff[r+1][i]--
	}

	for i := 1; i <= n; i++ {
		for j := range 3 {
			diff[i][j] += diff[i-1][j]
		}
	}

	dp := make(mat, 3)
	for i := range 3 {
		dp[i] = make([]int, 1)
	}
	dp[1][0] = 1

	getMat := func(c int) mat {
		res := make(mat, 3)
		for i := range 3 {
			res[i] = make([]int, 3)
		}
		for i := range 3 {
			for j := range 3 {
				if abs(i-j) <= 1 {
					res[i][j] = 1
				}
			}
		}
		for i := range 3 {
			if diff[c][i] > 0 {
				clear(res[i])
			}
		}
		return res
	}

	for i := 1; i < len(all); i++ {
		tr := getMat(i)
		dp = matPow(tr, all[i]-all[i-1]).mul(dp)
	}

	return dp[1][0]
}

const mod = 1e9 + 7

func add(nums ...int) int {
	var res int
	for _, num := range nums {
		res += num
		if res >= mod {
			res -= mod
		}
	}
	return res
}

func mul(a, b int) int {
	return a * b % mod
}

type mat [][]int

func (this mat) mul(that mat) mat {
	n := len(this)
	m := len(this[0])
	k := len(that[0])

	res := make(mat, n)
	for i := range n {
		res[i] = make([]int, k)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for u := 0; u < m; u++ {
				res[i][j] = add(res[i][j], mul(this[i][u], that[u][j]))
			}
		}
	}
	return res
}

func identityMat(n int) mat {
	res := make(mat, n)
	for i := range n {
		res[i] = make([]int, n)
		res[i][i] = 1
	}
	return res
}

func matPow(a mat, b int) mat {
	// n == m
	res := identityMat(len(a))

	for b > 0 {
		if b&1 == 1 {
			res = res.mul(a)
		}
		a = a.mul(a)
		b >>= 1
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
