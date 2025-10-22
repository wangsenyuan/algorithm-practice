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
	if res {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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

func drive(reader *bufio.Reader) bool {
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

var dd = []int{-1, 0, 1, 0, -1}

func solve(a []string) bool {
	n := len(a)
	m := len(a[0])

	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, m)
	}

	var dfs func(r int, c int, in int) bool
	dfs = func(r int, c int, in int) bool {
		if marked[r][c] {
			return true
		}
		marked[r][c] = true

		for i := range 4 {
			if in >= 0 && i != in && i&1 == in&1 {
				continue
			}
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && a[nr][nc] == a[r][c] && dfs(nr, nc, i) {
				return true
			}
		}
		return false
	}

	for i := range n {
		for j := range m {
			if !marked[i][j] {
				if dfs(i, j, -1) {
					return true
				}
			}
		}
	}
	return false
}
