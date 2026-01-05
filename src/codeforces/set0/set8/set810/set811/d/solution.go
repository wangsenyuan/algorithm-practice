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
	n, _ := readTwoNums(reader)
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}

	ask := func(cmd string) []int {
		fmt.Println(cmd)
		res := readNNums(reader, 2)
		res[0]--
		res[1]--
		return res
	}
	solve(a, ask)
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

var dd = []int{-1, 0, 1, 0, -1}

func solve(a []string, play func(string) []int) {
	n := len(a)
	m := len(a[0])

	dist := make([][]int, n)
	fa := make([][]int, n)

	for i := range n {
		dist[i] = make([]int, m)
		fa[i] = make([]int, m)
		for j := range m {
			dist[i][j] = -1
		}
	}

	que := make([]int, n*m)
	var head, tail int
	que[head] = 0
	head++
	dist[0][0] = 0

	dest := []int{-1, 01}

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++

		if a[r][c] == 'F' {
			dest = []int{r, c}
		}

		for i := range 4 {
			nr, nc := r+dd[i], c+dd[i+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < m && dist[nr][nc] == -1 && a[nr][nc] != '*' {
				dist[nr][nc] = dist[r][c] + 1
				que[head] = nr*m + nc
				head++
				fa[nr][nc] = r*m + c
			}
		}
	}

	var path [][]int
	for r, c := dest[0], dest[1]; r != 0 || c != 0; r, c = fa[r][c]/m, fa[r][c]%m {
		path = append(path, []int{r, c})
	}
	path = append(path, []int{0, 0})
	slices.Reverse(path)

	flag := []int{0, 0}

	for i := 1; i < len(path); i++ {
		r, c := path[i-1][0], path[i-1][1]
		nr, nc := path[i][0], path[i][1]
		if r == nr {
			// 同一行
			if flag[0] == 0 {
				// 水平方向还没有被决定
				cmd := "L"
				if c < nc {
					cmd = "R"
				}
				res := play(cmd)
				if res[1] == c {
					flag[0] = -1
					// pos should not change
					i--
					continue
				} else {
					flag[0] = 1
				}
			} else {
				cmd := "L"
				if c < nc {
					cmd = "R"
				}
				if flag[0] == -1 {
					cmd = reverse(cmd)
				}
				play(cmd)
			}
		} else {
			// c == nc
			if flag[1] == 0 {
				cmd := "U"
				if r < nr {
					cmd = "D"
				}
				res := play(cmd)
				if res[0] == r {
					flag[1] = -1
					i--
					continue
				} else {
					flag[1] = 1
				}
			} else {
				cmd := "U"
				if r < nr {
					cmd = "D"
				}
				if flag[1] == -1 {
					cmd = reverse(cmd)
				}
				play(cmd)
			}
		}
	}
}

func reverse(cmd string) string {
	if cmd == "L" {
		return "R"
	}
	if cmd == "R" {
		return "L"
	}
	if cmd == "U" {
		return "D"
	}
	return "U"
}
