package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1])
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

func drive(reader *bufio.Reader) []int {
	n, _ := readTwoNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)
	}
	return solve(grid)
}

func solve(grid []string) []int {
	n := len(grid)
	m := len(grid[0])

	prev := make([][][2]int, n)
	next := make([][][2]int, n)
	for i := range n {
		prev[i] = make([][2]int, m)
		next[i] = make([][2]int, m)
	}

	reset := func(x int, y int) (int, int) {
		nx, ny := x, y
		switch grid[x][y] {
		case 'U':
			nx = prev[x][y][1]
		case 'D':
			nx = next[x][y][1]
		case 'L':
			ny = prev[x][y][0]
		default:
			ny = next[x][y][0]
		}

		l, t, r, b := prev[x][y][0], prev[x][y][1], next[x][y][0], next[x][y][1]
		if l != -1 {
			next[x][l][0] = r
		}
		if r < m {
			prev[x][r][0] = l
		}
		if t != -1 {
			next[t][y][1] = b
		}
		if b < n {
			prev[b][y][1] = t
		}

		return nx, ny
	}

	check := func(x int, y int) int {
		for i := range n {
			last := -1
			for j := range m {
				prev[i][j][0] = last
				if grid[i][j] != '.' {
					last = j
				}
			}
			last = m
			for j := m - 1; j >= 0; j-- {
				next[i][j][0] = last
				if grid[i][j] != '.' {
					last = j
				}
			}
		}
		for j := range m {
			last := -1
			for i := range n {
				prev[i][j][1] = last
				if grid[i][j] != '.' {
					last = i
				}
			}
			last = n
			for i := n - 1; i >= 0; i-- {
				next[i][j][1] = last
				if grid[i][j] != '.' {
					last = i
				}
			}
		}

		var score int

		for {
			score++
			x, y = reset(x, y)
			if x < 0 || x >= n || y < 0 || y >= m {
				break
			}
		}

		return score
	}
	ans := []int{0, 0}

	for i := range n {
		for j := range m {
			if grid[i][j] != '.' {
				tmp := check(i, j)
				if tmp > ans[0] {
					ans[0] = tmp
					ans[1] = 1
				} else if tmp == ans[0] {
					ans[1]++
				}
			}
		}
	}

	return ans
}
