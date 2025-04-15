package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
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
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = readNNums(reader, n)
	}
	return solve(n, m, grid)
}

const inf = 1 << 60

func solve(n int, m int, grid [][]int) int {

	dp := make([]int, n)
	gp := make([]int, n)
	col := make([]int, n)
	fp := make([]int, n)
	sp := make([]int, n)
	check := func() int {
		clear(col)
		for i := 0; i < n; i++ {
			dp[i] = -inf
			fp[i] = -inf
			sp[i] = -inf
			gp[i] = -inf
		}

		for i := 0; i < n; i++ {
			if i >= m {
				for j := 0; j < n; j++ {
					col[j] -= grid[i-m][j]
				}
			}
			var sum int
			for j := 0; j < n; j++ {
				col[j] += grid[i][j]
				sum += col[j]
				if j >= m {
					sum -= col[j-m]
				}
				if j >= m {
					dp[i] = max(dp[i], sum+fp[j-m])
				}
				if i >= m {
					dp[i] = max(dp[i], sum+gp[i-m])
				}

				if j >= m-1 {
					fp[j] = max(fp[j], sum)
					if j > 0 {
						fp[j] = max(fp[j], fp[j-1])
					}
				}
				gp[i] = max(gp[i], fp[j])
			}
			sum = 0
			for j := n - 1; j >= 0; j-- {
				sum += col[j]
				if j+m < n {
					sum -= col[j+m]
				}
				if j+m < n {
					dp[i] = max(dp[i], sum+sp[j+m])
				}
				if j+m <= n && i >= m-1 {
					sp[j] = max(sp[j], sum)
					if j+1 < n {
						sp[j] = max(sp[j], sp[j+1])
					}
				}
			}
			if i > 0 {
				dp[i] = max(dp[i], dp[i-1])
				gp[i] = max(gp[i], gp[i-1])
			}
		}

		clear(col)
		var best int
		for i := n - 1; i >= 0; i-- {
			if n-i > m {
				for j := 0; j < n; j++ {
					col[j] -= grid[i+m][j]
				}
			}
			var sum int
			for j := 0; j < n; j++ {
				col[j] += grid[i][j]
				sum += col[j]
				if j >= m {
					sum -= col[j-m]
				}
				if j >= m-1 && i+m <= n && i >= m {
					best = max(best, sum+dp[i-1])
				}
			}
		}
		return best
	}
	buf := make([][]int, n)
	for i := 0; i < n; i++ {
		buf[i] = make([]int, n)
	}
	rotate := func() {
		// 旋转90度
		for i := 0; i < n; i++ {
			copy(buf[i], grid[i])
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ni := j
				nj := n - 1 - i
				grid[i][j] = buf[ni][nj]
			}
		}
	}

	var res int
	for i := 0; i < 4; i++ {
		res = max(res, check())
		rotate()
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func clear(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = 0
	}
}
