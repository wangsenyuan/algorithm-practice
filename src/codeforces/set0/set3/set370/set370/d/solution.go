package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Fprintln(writer, -1)
		return
	}
	for _, row := range res {
		fmt.Fprintln(writer, row)
	}
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

func drive(reader *bufio.Reader) (a []string, d int, res []string) {
	n, _ := readTwoNums(reader)
	a = make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = readString(reader)
	}
	d, res = solve(a)
	return
}

type pair struct {
	first  int
	second int
}

func solve(a []string) (int, []string) {
	n := len(a)
	m := len(a[0])
	var arr []pair
	for i := range n {
		for j := range m {
			if a[i][j] == 'w' {
				arr = append(arr, pair{i, j})
			}
		}
	}

	if len(arr) == 1 {
		return 1, a
	}
	var last_row, last_col int
	first_row, first_col := n, m
	for _, cur := range arr {
		first_row = min(first_row, cur.first)
		first_col = min(first_col, cur.second)
		last_row = max(last_row, cur.first)
		last_col = max(last_col, cur.second)
	}
	d := max(last_row-first_row+1, last_col-first_col+1)

	if d > min(n, m) || len(arr) > 4*(d-1) {
		return -1, nil
	}
	if d == last_row-first_row+1 {
		l := max(0, last_col-d+1)
		for j := l; j+d <= m && j <= first_col; j++ {
			if check(first_row, j, last_row, j+d-1, arr) {
				return d, construct(n, m, arr, d, first_row, j)
			}
		}
	}

	if d == last_col-first_col+1 {
		t := max(0, last_row-d+1)
		for i := t; i+d <= n && i <= first_row; i++ {
			if check(i, first_col, i+d-1, last_col, arr) {
				return d, construct(n, m, arr, d, i, first_col)
			}
		}
	}

	return -1, nil
}

func check(x1 int, y1 int, x2 int, y2 int, arr []pair) bool {
	for _, cur := range arr {
		if cur.first == x1 && y1 <= cur.second && cur.second <= y2 {
			// top
			continue
		}
		if cur.first == x2 && y1 <= cur.second && cur.second <= y2 {
			// bottom
			continue
		}

		if cur.second == y1 && x1 <= cur.first && cur.first <= x2 {
			// left
			continue
		}
		if cur.second == y2 && x1 <= cur.first && cur.first <= x2 {
			// right
			continue
		}
		return false
	}
	return true
}

func construct(n int, m int, arr []pair, d int, x int, y int) []string {
	buf := make([][]byte, n)
	for i := range n {
		buf[i] = make([]byte, m)
		for j := range m {
			buf[i][j] = '.'
		}
	}
	for j := y; j < y+d; j++ {
		buf[x][j] = '+'
		buf[x+d-1][j] = '+'
	}
	for i := x; i < x+d; i++ {
		buf[i][y] = '+'
		buf[i][y+d-1] = '+'
	}
	for _, cur := range arr {
		buf[cur.first][cur.second] = 'w'
	}
	res := make([]string, n)
	for i := range n {
		res[i] = string(buf[i])
	}
	return res
}
