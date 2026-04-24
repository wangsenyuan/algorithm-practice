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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	tc := readNums(reader)[0]
	for range tc {
		_, res := drive(reader)
		if len(res) == 0 {
			fmt.Fprintln(writer, "No")
		} else {
			fmt.Fprintln(writer, "Yes")
			for _, e := range res {
				fmt.Fprintln(writer, e[0], e[1])
			}
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)

	ss := strings.Split(s, " ")
	res := make([]int, len(ss))
	for i := range len(ss) {
		res[i], _ = strconv.Atoi(ss[i])
	}
	return res
}

func drive(reader *bufio.Reader) (s []string, res [][]int) {
	n := readNums(reader)[0]
	s = make([]string, n)
	for i := range n {
		s[i] = readString(reader)
	}
	res = solve(s)
	return
}

func solve(s []string) [][]int {
	n := len(s)

	var res [][]int

	findMid := func(x int, y int) int {
		for u := range n {
			if x == u || y == u {
				continue
			}
			if s[x][u] == '1' && s[u][y] == '1' {
				return u
			}
		}
		return -1
	}

	for i := range n {
		if s[i][i] == '0' {
			return nil
		}
		for j := range n {
			if i != j && s[i][j] == '1' {
				u := findMid(i, j)
				if u < 0 {
					res = append(res, []int{i + 1, j + 1})
				}
			}
		}
	}

	if len(res) != n-1 || !verify(s, res) {
		return nil
	}

	return res
}

func verify(s []string, edges [][]int) bool {
	n := len(s)

	fa := make([]int, n)
	for i := range n {
		fa[i] = i
	}
	find := func(x int) int {
		y := x
		for fa[y] != y {
			y = fa[y]
		}
		for fa[x] != y {
			fa[x], x = y, fa[x]
		}
		return y
	}

	sz := n
	union := func(x, y int) {
		x = find(x)
		y = find(y)
		if x == y {
			return
		}
		fa[x] = y
		sz--
	}

	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0]-1, e[1]-1
		adj[u] = append(adj[u], v)
		union(u, v)
	}

	que := make([]int, n)
	check := func(x int) string {
		buf := make([]byte, n)
		for i := range n {
			buf[i] = '0'
		}
		buf[x] = '1'
		var head, tail int
		que[head] = x
		head++
		for tail < head {
			u := que[tail]
			tail++
			for _, v := range adj[u] {
				if buf[v] == '0' {
					buf[v] = '1'
					que[head] = v
					head++
				}
			}
		}
		return string(buf)
	}

	for i := range n {
		tmp := check(i)
		if s[i] != tmp {
			return false
		}
	}
	return sz == 1
}
