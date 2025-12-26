package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, ans := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, res)
	for _, s := range ans {
		fmt.Fprintln(writer, s)
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

func drive(reader *bufio.Reader) (int, []string) {
	n, _, k := readThreeNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(k, a)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(k int, a []string) (int, []string) {
	n := len(a)
	m := len(a[0])

	marked := make([][]int, n)

	for i := range n {
		marked[i] = make([]int, m)
	}

	var dfs func(r int, c int) int
	dfs = func(r int, c int) int {
		marked[r][c] = 1

		sz := 1
		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && a[nr][nc] == '.' && marked[nr][nc] == 0 {
				sz += dfs(nr, nc)
			}
		}
		return sz
	}

	for i := range n {
		if a[i][0] == '.' && marked[i][0] == 0 {
			dfs(i, 0)
		}
		if a[i][m-1] == '.' && marked[i][m-1] == 0 {
			dfs(i, m-1)
		}
	}

	for j := range m {
		if a[0][j] == '.' && marked[0][j] == 0 {
			dfs(0, j)
		}
		if a[n-1][j] == '.' && marked[n-1][j] == 0 {
			dfs(n-1, j)
		}
	}

	type lake struct {
		r  int
		c  int
		sz int
	}

	var lakes []lake

	for i := range n {
		for j := range m {
			if marked[i][j] == 0 && a[i][j] == '.' {
				sz := dfs(i, j)
				lakes = append(lakes, lake{i, j, sz})
			}
		}
	}

	slices.SortFunc(lakes, func(a, b lake) int {
		return a.sz - b.sz
	})

	buf := make([][]byte, n)
	for i := range n {
		buf[i] = []byte(a[i])
	}

	var dfs2 func(r int, c int)
	dfs2 = func(r int, c int) {
		buf[r][c] = '*'
		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && buf[nr][nc] == '.' {
				dfs2(nr, nc)
			}
		}
	}

	var res int
	for i := range len(lakes) - k {
		res += lakes[i].sz
		dfs2(lakes[i].r, lakes[i].c)
	}

	ans := make([]string, n)
	for i := range n {
		ans[i] = string(buf[i])
	}

	return res, ans
}
