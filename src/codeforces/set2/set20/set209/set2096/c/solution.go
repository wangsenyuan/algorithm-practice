package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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
	n := readNum(reader)
	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = readNNums(reader, n)
	}
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(grid, a, b)
}

func solve(grid [][]int, a []int, b []int) int {
	ans := calc(grid, b)
	n := len(grid)
	tmp := make([][]int, n)
	for i := range n {
		tmp[i] = make([]int, n)
	}
	for i := range n {
		for j := range n {
			tmp[j][i] = grid[i][j]
		}
	}
	ans += calc(tmp, a)
	if ans >= inf {
		return -1
	}
	return ans
}

const inf = 1 << 60

func calc(grid [][]int, a []int) int {
	n := len(grid)
	dp := []int{0, a[0]}

	for j := 1; j < n; j++ {
		var flag int
		for i := 0; i < n; i++ {
			if grid[i][j] == grid[i][j-1] {
				flag |= 1
			} else if grid[i][j] == grid[i][j-1]+1 {
				flag |= 2
			} else if grid[i][j]+1 == grid[i][j-1] {
				flag |= 4
			}
		}
		if flag == 7 {
			return inf
		}
		fp := []int{inf, inf}
		if flag&1 == 1 {
			// 必须改变其中的一个
			if flag&4 == 0 {
				// 改变j，不改变j-1
				fp[1] = min(fp[1], dp[0]+a[j])
			}

			if flag&2 == 0 {
				// 改变j-1
				fp[0] = min(fp[0], dp[1])
			}
		} else {
			// 没有限制只改一个
			// 两个同时不变
			fp[0] = dp[0]
			// 或者同时改变
			fp[1] = min(inf, dp[1]+a[j])
			if flag&2 == 0 {
				// 改变j-1时，可以不改变j
				fp[0] = min(fp[0], dp[1])
			}
			if flag&4 == 0 {
				// 改变j时，可以不改变j-1
				fp[1] = min(fp[1], dp[0]+a[j])
			}
		}
		copy(dp, fp)
	}

	return min(dp[0], dp[1])
}
