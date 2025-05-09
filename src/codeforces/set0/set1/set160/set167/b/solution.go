package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.9f\n", res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) float64 {
	n, l, k := readThreeNums(reader)
	p := readNNums(reader, n)
	a := readNNums(reader, n)
	return solve(l, k, p, a)
}

func solve(l int, k int, p []int, a []int) float64 {
	n := len(a)
	k = min(k, n)
	dp := make([][]float64, n+1)
	fp := make([][]float64, n+1)
	for i := range n + 1 {
		dp[i] = make([]float64, 2*n+3)
		fp[i] = make([]float64, 2*n+3)
	}
	dp[0][n+k] = 1

	for i := range n {
		f := float64(p[i]) / 100
		for u := 0; u <= i; u++ {
			for v := -n; v <= n; v++ {
				if dp[u][v+n] == 0 {
					continue
				}
				if a[i] == -1 {
					// 获胜的情况下，会消耗一个容量
					fp[u+1][v-1+n] += dp[u][v+n] * f
				} else {
					// 获胜的情况下，增加容量
					fp[u+1][min(v+a[i], n)+n] += dp[u][v+n] * f
				}
				fp[u][v+n] += dp[u][v+n] * (1 - f)
			}
		}
		for u := 0; u <= i+1; u++ {
			for v := -n; v <= n; v++ {
				dp[u][v+n] = fp[u][v+n]
				fp[u][v+n] = 0
			}
		}
	}
	var ans float64
	for u := l; u <= n; u++ {
		for v := 0; v <= n; v++ {
			ans += dp[u][v+n]
		}
	}
	return ans
}
