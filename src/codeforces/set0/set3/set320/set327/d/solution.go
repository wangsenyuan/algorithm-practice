package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(cur)
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func process(reader *bufio.Reader) ([]string, []string) {
	n, m := readTwoNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)[:m]
	}

	return solve(grid), grid
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(grid []string) []string {

	n := len(grid)
	m := len(grid[0])

	buf := make([][]byte, n)
	for i := range n {
		buf[i] = []byte(grid[i])
	}

	var res []string
	var dfs func(p int, u int)
	dfs = func(p int, u int) {
		r, c := u/m, u%m
		res = append(res, fmt.Sprintf("B %d %d", r+1, c+1))
		buf[r][c] = '#'
		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < n && y >= 0 && y < m && buf[x][y] == '.' {
				dfs(u, x*m+y)
			}
		}
		if p >= 0 {
			res = append(res, fmt.Sprintf("D %d %d", r+1, c+1))
			res = append(res, fmt.Sprintf("R %d %d", r+1, c+1))
		}
	}

	for i := range n {
		for j := range m {
			if buf[i][j] == '.' {
				dfs(-1, i*m+j)
			}
		}
	}

	return res
}
