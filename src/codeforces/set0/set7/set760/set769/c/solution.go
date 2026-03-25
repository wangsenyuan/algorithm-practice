package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	nums := make([]int, len(ss))
	for i := range nums {
		nums[i], _ = strconv.Atoi(ss[i])
	}
	return nums
}

func drive(reader *bufio.Reader) string {
	nums := readNums(reader)
	n := nums[0]
	k := nums[2]
	maze := make([]string, n)
	for i := range n {
		maze[i] = readString(reader)
	}
	return solve(k, maze)
}

func solve(k int, maze []string) string {
	if k&1 == 1 {
		return "IMPOSSIBLE"
	}
	n := len(maze)
	m := len(maze[0])
	sx, sy := 0, 0

	dist := make([][]int, n)

	for i := range n {
		dist[i] = make([]int, m)
		for j := range m {
			dist[i][j] = -1
			if maze[i][j] == 'X' {
				sx, sy = i, j
			}
		}
	}
	que := make([]int, n*m)
	var head, tail int
	que[head] = sx*m + sy
	head++
	dist[sx][sy] = 0

	dirs := "DLRU"

	dd := [][]int{
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, 0},
	}
	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++
		for i := range 4 {
			x, y := r+dd[i][0], c+dd[i][1]
			if x >= 0 && x < n && y >= 0 && y < m && dist[x][y] < 0 && maze[x][y] != '*' {
				dist[x][y] = dist[r][c] + 1
				que[head] = x*m + y
				head++
			}
		}
	}

	buf := make([]byte, k)

	r, c := sx, sy

found:
	for d := range k {
		for i := range 4 {
			nr, nc := r+dd[i][0], c+dd[i][1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && maze[nr][nc] != '*' && dist[nr][nc] <= k-d-1 {
				buf[d] = dirs[i]
				r, c = nr, nc
				continue found
			}
		}
		return "IMPOSSIBLE"
	}

	return string(buf)
}
