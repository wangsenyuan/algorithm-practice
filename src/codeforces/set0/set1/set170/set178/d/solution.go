package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	sum, ans := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, sum)
	for _, row := range ans {
		for _, v := range row {
			fmt.Fprint(writer, v, " ")
		}
		fmt.Fprintln(writer)
	}
}

func drive(reader *bufio.Reader) (sum int, ans [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n*n)
	for i := range n * n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

func solve(n int, a []int) (sum int, ans [][]int) {
	rand.Shuffle(len(a), func(i, j int) {
		a[i], a[j] = a[j], a[i]
	})
	for _, v := range a {
		sum += v
	}

	sum /= n

	ans = make([][]int, n)
	for i := range n {
		ans[i] = make([]int, n)
	}

	getRow := func(r int) int {
		var res int
		for c := range n {
			res += ans[r][c]
		}
		return res
	}

	getCol := func(c int) int {
		var res int
		for r := range n {
			res += ans[r][c]
		}
		return res
	}

	var f1 func(i int, mask int) bool
	var f2 func(r int, c int, mask int) bool
	var f3 func(r int, c int, mask int) bool

	f1 = func(i int, mask int) bool {
		if i == n {
			// check diagonal
			var s1 int
			var s2 int
			for r := range n {
				s1 += ans[r][r]
				s2 += ans[r][n-1-r]
			}
			return s1 == sum && s2 == sum
		}
		return f2(i, i, mask)
	}

	f2 = func(r int, c int, mask int) bool {
		// 从位置 (i, i) 开始放置
		if c == n {
			if getRow(r) != sum {
				return false
			}
			return f3(r+1, r, mask)
		}
		// c < n
		for i, v := range a {
			if (mask>>i)&1 == 0 {
				ans[r][c] = v
				if f2(r, c+1, mask|(1<<i)) {
					return true
				}
			}
		}
		return false
	}

	f3 = func(r int, c int, mask int) bool {
		if r == n {
			if getCol(c) != sum {
				return false
			}
			return f1(c+1, mask)
		}

		for i, v := range a {
			if (mask>>i)&1 == 0 {
				ans[r][c] = v
				if f3(r+1, c, mask|(1<<i)) {
					return true
				}
			}
		}
		return false
	}

	f1(0, 0)

	return
}
