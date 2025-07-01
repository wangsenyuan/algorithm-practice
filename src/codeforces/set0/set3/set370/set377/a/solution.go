package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(x)
		buf.WriteByte('\n')
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (k int, maze []string, res []string) {
	n, m, k := readThreeNums(reader)
	maze = make([]string, n)
	for i := 0; i < n; i++ {
		maze[i] = readString(reader)[:m]
	}
	res = solve(k, maze)
	return
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(k int, maze []string) []string {

	n := len(maze)
	m := len(maze[0])

	var s int
	marked := make([][]bool, n)
	for i := 0; i < n; i++ {
		marked[i] = make([]bool, m)
		for j := range m {
			if maze[i][j] == '.' {
				s++
			}
		}
	}

	s -= k

	var dfs func(r int, c int)
	dfs = func(r int, c int) {
		if s == 0 {
			return
		}
		s--
		marked[r][c] = true
		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && maze[x][y] == '.' && !marked[x][y] {
				dfs(x, y)
			}
		}
	}

	buf := make([][]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = []byte(maze[i])
		for j := 0; j < m; j++ {
			if maze[i][j] == '.' {
				dfs(i, j)
			}
		}
	}

	ans := make([]string, n)
	for i := range n {
		for j := range m {
			if buf[i][j] == '.' && !marked[i][j] {
				buf[i][j] = 'X'
			}
		}
		ans[i] = string(buf[i])
	}

	return ans
}
