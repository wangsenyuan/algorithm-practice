package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)
	grid := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &grid[i])
	}
	targets := make([]int, q)
	for i := range q {
		fmt.Fscan(reader, &targets[i])
	}
	return solve(grid, targets)
}

const inf = 1 << 60

func solve(grid []string, targets []int) []int {
	// TODO: implement
	// assume n <= m => n < 200 => 如果是竖向的和斜向的表达式, 长度不会超过n
	// 但是横向的,还是会很长, 但最长也就是1000, 所以,似乎也没有问题, 3*1e7
	// n * m <= 3 * 1e4, *的优先级高于
	// 单独数字组成的比较好处理(只需要横竖斜处理6格就能全部找出来)
	// 问题出在表达式上面(因为表达式可以一直增加, 比如 1*1*1*1*1...)
	dp := make([][]int, 4)
	for i := range 4 {
		dp[i] = solve1(grid)
		grid = rotate(grid)
	}

	dp1 := make([]int, 1e6+1)
	for _, row := range grid {
		for _, c := range row {
			if c >= '1' && c <= '9' {
				dp1[int(c-'0')]++
			}
		}
	}

	ans := make([]int, len(targets))

	for i, v := range targets {
		for j := range dp {
			ans[i] += dp[j][v]
		}
		ans[i] += dp1[v]
	}

	return ans
}

func rotate(grid []string) []string {
	n, m := len(grid), len(grid[0])
	buf := make([][]byte, m)
	for i := range m {
		buf[i] = make([]byte, n)
	}
	for i := range n {
		for j := range m {
			buf[j][i] = grid[n-1-i][j]
		}
	}
	ans := make([]string, m)
	for i := range m {
		ans[i] = string(buf[i])
	}
	return ans
}

func solve1(grid []string) []int {
	n, m := len(grid), len(grid[0])

	dp := make([]int, 1e6+10)

	stack := make([]int, 3)
	op := make([]byte, 3)

	calc := func(num1 int, op byte, num2 int) int {
		if op == '+' {
			return num1 + num2
		}
		return num1 * num2
	}

	var top int
	play := func(r0 int, c0 int, r int, c int, dr int, dc int) bool {
		if grid[r][c] == '+' || grid[r][c] == '*' {
			if c == c0 || (grid[r+dr][c+dc] == '+' || grid[r+dr][c+dc] == '*') {
				// invalid
				return false
			}
			for top > 1 && prior(op[top-2], grid[r][c]) {
				stack[top-2] = calc(stack[top-2], op[top-2], stack[top-1])
				top--
			}
			if stack[top-1] > 1e6 {
				return false
			}
			op[top-1] = grid[r][c]
		} else {
			v := int(grid[r][c] - '0')
			if c == c0 || grid[r+dr][c+dc] == '+' || grid[r+dr][c+dc] == '*' {
				stack[top] = v
				top++
			} else {
				stack[top-1] = stack[top-1]*10 + v
			}
			if top == 1 && c > c0 {
				// 单个格子的情况,单独处理
				if stack[top-1] <= 1e6 {
					dp[stack[top-1]]++
				}
			} else if top == 2 {
				tmp := calc(stack[top-2], op[top-2], stack[top-1])
				if tmp <= 1e6 {
					dp[tmp]++
				}
			} else if top == 3 {
				// a + b * c (只有这种情况)?
				tmp := stack[top-3] + stack[top-2]*stack[top-1]
				if tmp <= 1e6 {
					dp[tmp]++
				}
			}

			if stack[top-1] > 1e6 {
				return false
			}
		}
		return true
	}

	// 先处理横向
	for i := range n {
		for j1 := range m {
			top = 0
			for j := j1; j < m; j++ {
				ok := play(i, j1, i, j, 0, -1)
				if !ok {
					break
				}
			}

			top = 0
			// 现在处理斜边的情况
			for d := 0; j1+d < m && i+d < n; d++ {
				r, c := i+d, j1+d
				ok := play(i, j1, r, c, -1, -1)
				if !ok {
					break
				}
			}

		}
	}

	return dp
}

func prior(a byte, b byte) bool {
	return a == '*' || b == '+'
}
