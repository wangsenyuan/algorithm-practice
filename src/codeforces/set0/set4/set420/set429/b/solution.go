package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, m)
	}
	return solve(a)
}

func solve(a [][]int) int {
	n := len(a)
	m := len(a[0])
	dp := make([][][]int, 4)
	for i := range 4 {
		dp[i] = make([][]int, n)
		for r := range n {
			dp[i][r] = make([]int, m)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dp[0][i][j] = a[i][j]
			if i > 0 && j > 0 {
				dp[0][i][j] += max(dp[0][i-1][j], dp[0][i][j-1])
			} else if i > 0 {
				dp[0][i][j] += dp[0][i-1][j]
			} else if j > 0 {
				dp[0][i][j] += dp[0][i][j-1]
			}
		}
		for j := m - 1; j >= 0; j-- {
			dp[1][i][j] = a[i][j]
			if i > 0 && j < m-1 {
				dp[1][i][j] += max(dp[1][i-1][j], dp[1][i][j+1])
			} else if i > 0 {
				dp[1][i][j] += dp[1][i-1][j]
			} else if j < m-1 {
				dp[1][i][j] += dp[1][i][j+1]
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			dp[2][i][j] = a[i][j]
			if i < n-1 && j < m-1 {
				dp[2][i][j] += max(dp[2][i+1][j], dp[2][i][j+1])
			} else if i < n-1 {
				dp[2][i][j] += dp[2][i+1][j]
			} else if j < m-1 {
				dp[2][i][j] += dp[2][i][j+1]
			}
		}
		for j := range m {
			dp[3][i][j] = a[i][j]
			if i < n-1 && j > 0 {
				dp[3][i][j] += max(dp[3][i+1][j], dp[3][i][j-1])
			} else if i < n-1 {
				dp[3][i][j] += dp[3][i+1][j]
			} else if j > 0 {
				dp[3][i][j] += dp[3][i][j-1]
			}
		}
	}
	var best int
	for i := range n {
		for j := range m {
			if i == 0 || j == 0 || i == n-1 || j == m-1 {
				// 无法在边上相遇，不满足只在一个cell里面相遇的条件
				continue
			}
			tmp1 := dp[0][i-1][j] + dp[3][i][j-1] + dp[1][i][j+1] + dp[2][i+1][j]
			best = max(best, tmp1)
			tmp2 := dp[0][i][j-1] + dp[3][i+1][j] + dp[1][i-1][j] + dp[2][i][j+1]
			best = max(best, tmp2)
		}
	}
	return best
}
