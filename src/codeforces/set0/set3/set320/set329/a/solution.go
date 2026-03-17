package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, ok, ans := drive(reader)
	if !ok {
		fmt.Println("-1")
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, cur := range ans {
		fmt.Fprintln(writer, cur[0], cur[1])
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

func drive(reader *bufio.Reader) (n int, a []string, ok bool, ans [][]int) {
	n = readNums(reader)[0]
	a = make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	ok, ans = solve(n, a)
	return
}

func solve(n int, grid []string) (ok bool, ans [][]int) {
	row := make([]int, n)
	col := make([]int, n)
	for i := range n {
		for j := range n {
			if grid[i][j] == 'E' {
				row[i]++
				col[j]++
			}
		}
	}
	w := slices.Max(row)
	v := slices.Max(col)
	if min(w, v) == n {
		return false, nil
	}
	if w < n {
		// 每行都存在一个空位
		for i := range n {
			for j := range n {
				if grid[i][j] == '.' {
					ans = append(ans, []int{i + 1, j + 1})
					break
				}
			}
		}
	} else {
		// v < n
		for j := range n {
			for i := range n {
				if grid[i][j] == '.' {
					ans = append(ans, []int{i + 1, j + 1})
					break
				}
			}
		}
	}

	return true, ans
}
