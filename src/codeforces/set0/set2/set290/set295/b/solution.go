package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, n)
	}
	x := readNNums(reader, n)
	return solve(a, x)
}

const inf = 1 << 60

func solve(a [][]int, x []int) []int {
	n := len(a)
	for i := range n {
		x[i]--
	}
	ans := make([]int, n)

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		for j := range n {
			dp[i][j] = inf
		}
	}

	f := func(pos int, u int) {
		for i := pos + 1; i < n; i++ {
			v := x[i]
			dp[u][v] = a[u][v]
		}
		for i := pos + 1; i < n; i++ {
			v := x[i]
			for j := pos + 1; j < n; j++ {
				w := x[j]
				dp[u][v] = min(dp[u][v], dp[u][w]+dp[w][v])
			}
		}
	}

	g := func(pos int, u int) {
		for i := pos + 1; i < n; i++ {
			r := x[i]
			dp[r][u] = a[r][u]
			for j := pos + 1; j < n; j++ {
				if i == j {
					continue
				}
				v := x[j]
				dp[r][u] = min(dp[r][u], dp[r][v]+a[v][u])
			}
			for j := pos + 1; j < n; j++ {
				if i == j {
					continue
				}
				v := x[j]
				dp[r][v] = min(dp[r][v], dp[r][u]+dp[u][v])
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		u := x[i]
		// 先计算dp[u][v]
		f(i, u)
		g(i, u)
		for j := i; j < n; j++ {
			for k := j + 1; k < n; k++ {
				r := x[j]
				w := x[k]
				ans[i] += dp[r][w]
				ans[i] += dp[w][r]
			}
		}
	}

	return ans
}
