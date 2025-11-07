package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var k int
	fmt.Fscan(reader, &k)
	res := solve(k)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, s := range res {
		fmt.Fprintln(writer, string(s))
	}
}

func solve(k int) []string {
	if k == 0 {
		return []string{"+"}
	}

	res := make([][]byte, 1<<k)

	for i := range res {
		res[i] = make([]byte, 1<<k)
	}

	//上半部分 * 下半部分 = 0
	//那么 上部分和下半部分，必须有一半是相同的，另外一半是不同的
	// 如果 k = 1

	reverse := func(b byte) byte {
		if b == '+' {
			return '*'
		}
		return '+'
	}

	var f func(r int, c int, k int)
	f = func(r int, c int, k int) {
		if k == 1 {
			res[r][c] = '+'
			res[r][c+1] = '+'
			res[r+1][c] = '+'
			res[r+1][c+1] = '*'
			return
		}
		f(r, c, k-1)
		r1 := r + (1 << (k - 1))
		r2 := r + (1 << k)
		c1 := c + (1 << (k - 1))
		c2 := c + (1 << k)
		for i := r; i < r1; i++ {
			for j := c1; j < c2; j++ {
				res[i][j] = reverse(res[i][j-(1<<(k-1))])
			}
		}
		for i := r1; i < r2; i++ {
			for j := c; j < c1; j++ {
				res[i][j] = (res[i-(1<<(k-1))][j])
			}
			for j := c1; j < c2; j++ {
				res[i][j] = reverse(res[i-(1<<(k-1))][j])
			}
		}
	}

	f(0, 0, k)

	ans := make([]string, 1<<k)
	for i := range ans {
		ans[i] = string(res[i])
	}
	return ans
}
