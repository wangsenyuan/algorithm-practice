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

func process(reader *bufio.Reader) string {
	n, m := readTwoNums(reader)
	grid := make([]string, n)
	for i := range n {
		grid[i] = readString(reader)[:m]
	}
	if solve(grid) {
		return "YES"
	}
	return "NO"
}

func solve(grid []string) bool {
	n := len(grid)
	m := len(grid[0])
	sum := make([][]int, n+1)
	var start, dest []int
	sum[0] = make([]int, m+1)
	for i := range n {
		sum[i+1] = make([]int, m+1)
		for j := range m {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j]
			if grid[i][j] == '*' {
				sum[i+1][j+1]++
			}
			if grid[i][j] == 'S' {
				start = []int{i, j}
			}
			if grid[i][j] == 'T' {
				dest = []int{i, j}
			}
		}
	}

	get := func(x1, y1, x2, y2 int) int {
		x1, x2 = min(x1, x2), max(x1, x2)
		y1, y2 = min(y1, y2), max(y1, y2)
		return sum[x2+1][y2+1] - sum[x1][y2+1] - sum[x2+1][y1] + sum[x1][y1]
	}

	for j := range m {
		if get(start[0], start[1], start[0], j) == 0 &&
			get(dest[0], dest[1], dest[0], j) == 0 &&
			get(start[0], j, dest[0], j) == 0 {
			return true
		}
	}

	for i := range n {
		if get(start[0], start[1], i, start[1]) == 0 &&
			get(dest[0], dest[1], i, dest[1]) == 0 &&
			get(i, start[1], i, dest[1]) == 0 {
			return true
		}
	}

	return false
}
