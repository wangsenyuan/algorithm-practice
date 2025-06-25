package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	n, m := readTwoNums(reader)
	grid := make([]string, 4*n+1)
	for i := range grid {
		grid[i] = readString(reader)
	}
	return solve(n, m, grid)
}

const mod = 1e9 + 7

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func solve(n int, m int, grid []string) int {
	// 4 * n + 1, 4 * m + 1
	// i % 4 = 0 的是边界， j % 4 = 0的是边界，中间的是card的形状
	// 共有10种形状
	// 0 1 4 5 (2, 3, 6) (2_, 3_, 6_)

	flag := []int{0, 0, 1, 1, 0, 0, 1, 2, 2, 2}

	parse := func(u int, v int) int {
		var cnt int
		for i := u; i < u+3; i++ {
			for j := v; j < v+3; j++ {
				if grid[i][j] == 'O' {
					cnt++
				}
			}
		}
		if flag[cnt] == 0 {
			return cnt
		}
		// flag[cnt] == 1
		if cnt == 2 || cnt == 3 {
			if grid[u][v] == 'O' {
				cnt += 5
			}
		} else {
			// cnt = 6
			if grid[u][v+1] == 'O' {
				cnt = 9
			}
		}
		return cnt
	}

	mat := make([][]int, n)
	for i := range n {
		mat[i] = make([]int, m)
	}

	for i := 0; i < 4*n; i += 4 {
		for j := 0; j < 4*m; j += 4 {
			mat[i/4][j/4] = parse(i+1, j+1)
		}
	}

	checkVertical := func(c int) bool {
		if n%2 == 1 {
			// 必须是偶数行
			return false
		}
		for i := range n {
			if flag[mat[i][c]] == 2 {
				return false
			}
		}
		return true
	}

	fp := make([][]int, n+1)
	for i := range n + 1 {
		fp[i] = make([]int, 2)
	}

	calc := func(c int) int {
		// c & c + 1
		fp[0][0] = 1

		for i := 0; i < n; i++ {
			clear(fp[i+1])

			if flag[mat[i][c]] != 1 && flag[mat[i][c+1]] != 1 {
				// 可以水平放置
				fp[i+1][1] = add(fp[i][0], fp[i][1])
			}
			if i >= 1 &&
				flag[mat[i-1][c]] != 2 && flag[mat[i][c]] != 2 &&
				flag[mat[i-1][c+1]] != 2 && flag[mat[i][c+1]] != 2 {
				// 可以放置垂直的两个
				fp[i+1][1] = add(fp[i+1][1], fp[i-1][1])
				fp[i+1][0] = add(fp[i+1][0], fp[i-1][0])
			}
		}

		return fp[n][1]
	}

	dp := make([]int, m+1)
	dp[0] = 1

	for j := 1; j <= m; j++ {
		if checkVertical(j - 1) {
			dp[j] = dp[j-1]
		}
		if j >= 2 {
			dp[j] = add(dp[j], mul(dp[j-2], calc(j-2)))
		}
	}

	return dp[m]
}
