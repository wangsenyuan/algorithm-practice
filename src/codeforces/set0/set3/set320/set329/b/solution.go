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
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)[:m]
	}
	return solve(grid)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(grid []string) int {
	n := len(grid)
	m := len(grid[0])

	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, m)
	}

	que := make([]int, n*m)

	bfs := func(x int, y int) {
		for i := range n {
			for j := range m {
				dist[i][j] = -1
			}
		}
		dist[x][y] = 0
		var head, tail int
		que[head] = x*m + y
		head++
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			for i := range 4 {
				u, v := r+dd[i], c+dd[i+1]
				if u >= 0 && u < n && v >= 0 && v < m && dist[u][v] == -1 && grid[u][v] != 'T' {
					dist[u][v] = dist[r][c] + 1
					que[head] = u*m + v
					head++
				}
			}
		}
	}

	var exit []int
	var player []int
	for i := range n {
		for j := range m {
			if grid[i][j] == 'E' {
				exit = []int{i, j}
			}
			if grid[i][j] == 'S' {
				player = []int{i, j}
			}
		}
	}
	bfs(exit[0], exit[1])

	// 似乎可以达到n*m的时间
	cnt := make([]int, n*m+1)

	for i := range n {
		for j := range m {
			if grid[i][j] >= '0' && grid[i][j] <= '9' && dist[i][j] > 0 {
				cnt[dist[i][j]] += int(grid[i][j] - '0')
			}
		}
	}

	for i := 1; i <= n*m; i++ {
		cnt[i] += cnt[i-1]
	}

	bfs(player[0], player[1])

	return cnt[dist[exit[0]][exit[1]]]
}
