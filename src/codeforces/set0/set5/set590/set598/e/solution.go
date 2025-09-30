package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for range tc {
		res := drive(reader)
		// fmt.Fprintln(writer, res)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(writer)
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	return solve(n, m, k)
}

const inf = 1e9

const N = 30
const M = 30
const K = 50

var dp [][][]int

func init() {
	dp = make([][][]int, N+1)
	for i := range N + 1 {
		dp[i] = make([][]int, M+1)
		for j := range M + 1 {
			dp[i][j] = make([]int, K+1)
			for u := range K + 1 {
				dp[i][j][u] = -1
			}
		}
	}

	var f func(h int, w int, exp int) int

	f = func(h int, w int, exp int) (res int) {

		if dp[h][w][exp] != -1 {
			return dp[h][w][exp]
		}
		// 水平切
		defer func() {
			dp[h][w][exp] = res
		}()
		if exp == 0 || exp == w*h {
			return 0
		}

		res = inf
		for i := 1; i < h; i++ {
			if i*w <= exp {
				res = min(res, f(h-i, w, exp-i*w)+w*w)
			} else {
				// exp < i * w
				res = min(res, f(i, w, exp)+w*w)
			}
		}
		for j := 1; j < w; j++ {
			if j*h <= exp {
				res = min(res, f(h, w-j, exp-j*h)+h*h)
			} else {
				res = min(res, f(h, j, exp)+h*h)
			}
		}
		return
	}

	for n := 1; n <= N; n++ {
		for m := 1; m <= M; m++ {
			for x := 1; x <= K && x <= N*M; x++ {
				f(n, m, x)
			}
		}
	}
}

func solve(n int, m int, k int) int {
	return dp[n][m][k]
}
