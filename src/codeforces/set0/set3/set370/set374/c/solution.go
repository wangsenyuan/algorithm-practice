package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
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

func process(reader *bufio.Reader) string {
	n, m := readTwoNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)[:m]
	}
	return solve(grid)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(grid []string) string {
	n := len(grid)
	m := len(grid[0])

	dp := make([][]int, n)
	marked := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
		marked[i] = make([]int, m)
	}

	var dima = "DIMA"

	var dfs func(r int, c int, expect int) (infinite bool, dist int)
	dfs = func(r int, c int, expect int) (infinite bool, dist int) {
		if marked[r][c] == 1 {
			return true, 0
		}
		if marked[r][c] == 2 {
			return false, dp[r][c]
		}

		expect++

		marked[r][c]++

		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == dima[expect%4] {
				in, dv := dfs(x, y, expect)

				if in {
					return true, 0
				}
				dp[r][c] = max(dp[r][c], dv)
			}
		}

		dp[r][c]++

		marked[r][c]++

		return false, dp[r][c]
	}
	var res int

	for r := range n {
		for c := range m {
			if grid[r][c] == 'D' && marked[r][c] == 0 {
				in, ds := dfs(r, c, 0)
				if in {
					return "Poor Inna!"
				}
				res = max(res, ds/4)
			}
		}
	}
	if res == 0 {
		return "Poor Dima!"
	}
	return fmt.Sprintf("%d", res)
}
