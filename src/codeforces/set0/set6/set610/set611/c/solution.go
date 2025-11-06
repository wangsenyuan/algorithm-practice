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
	res := drive(reader)
	for _, ans := range res {
		writer.WriteString(fmt.Sprintf("%d\n", ans))
	}
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
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	k := readNum(reader)
	queries := make([][]int, k)
	for i := range k {
		queries[i] = readNNums(reader, 4)
	}
	return solve(a, queries)
}

func solve(a []string, queries [][]int) []int {
	n := len(a)
	m := len(a[0])

	rows := make([][]int, n+1)
	cols := make([][]int, n+1)
	sum := make([][]int, n+1)
	for i := range n + 1 {
		sum[i] = make([]int, m+1)
		rows[i] = make([]int, m+1)
		cols[i] = make([]int, m+1)
	}

	get := func(i int, j int) int {
		if a[i][j] == '#' {
			return 0
		}
		var cnt int
		if i+1 < n && a[i+1][j] == '.' {
			cnt++
		}
		if j+1 < m && a[i][j+1] == '.' {
			cnt++
		}
		return cnt
	}
	for i := range n {
		for j := range m {
			sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + get(i, j)
			rows[i+1][j+1] = rows[i+1][j]
			if a[i][j] == '.' && j+1 < m && a[i][j+1] == '.' {
				rows[i+1][j+1]++
			}
			cols[i+1][j+1] = cols[i][j+1]
			if a[i][j] == '.' && i+1 < n && a[i+1][j] == '.' {
				cols[i+1][j+1]++
			}
		}
	}

	ask := func(r1 int, c1 int, r2 int, c2 int) int {
		if r1 == r2 {
			// 同一行
			return rows[r2][c2-1] - rows[r2][c1-1]
		}
		if c1 == c2 {
			return cols[r2-1][c2] - cols[r1-1][c2]
		}

		res := sum[r2-1][c2-1] - sum[r2-1][c1-1] - sum[r1-1][c2-1] + sum[r1-1][c1-1]
		res += rows[r2][c2-1] - rows[r2][c1-1]
		res += cols[r2-1][c2] - cols[r1-1][c2]
		return res
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		ans[i] = ask(cur[0], cur[1], cur[2], cur[3])
	}

	return ans
}
