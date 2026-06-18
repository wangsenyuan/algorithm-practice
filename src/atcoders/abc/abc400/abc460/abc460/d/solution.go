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
	res := drive(reader)
	for _, row := range res {
		fmt.Println(row)
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
	for i, v := range ss {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func drive(reader *bufio.Reader) []string {
	first := readNums(reader)
	n := first[0]
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) []string {
	n, m := len(a), len(a[0])

	mark := make([][]int, n)
	for i := range n {
		mark[i] = make([]int, m)
		for j := range m {
			mark[i][j] = -1
		}
	}

	hasWhiteNeighors := func(r int, c int) bool {
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < n && nc >= 0 && nc < m && a[nr][nc] == '.' {
					return true
				}
			}
		}
		return false
	}

	que := make([]int, n*m)
	var head, tail int
	for i := range n {
		for j := range m {
			if a[i][j] == '#' && hasWhiteNeighors(i, j) {
				mark[i][j] = 0
				que[head] = i*m + j
				head++
			}
		}
	}

	if head == 0 || head == n*m {
		return special(n, m)
	}

	for tail < head {
		r, c := que[tail]/m, que[tail]%m
		tail++
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < n && nc >= 0 && nc < m && mark[nr][nc] < 0 {
					mark[nr][nc] = 1 ^ mark[r][c]
					que[head] = nr*m + nc
					head++
				}
			}
		}
	}
	ans := make([]string, n)

	buf := make([]byte, m)
	for i := range n {
		for j := range m {
			buf[j] = '#'
			if mark[i][j] < 0 || mark[i][j]&1 == 1 {
				buf[j] = '.'
			}
		}
		ans[i] = string(buf)
	}

	return ans
}

func special(n int, m int) []string {
	res := make([]string, n)
	for i := range n {
		buf := make([]byte, m)
		for j := range m {
			buf[j] = '.'
		}
		res[i] = string(buf)
	}
	return res
}
