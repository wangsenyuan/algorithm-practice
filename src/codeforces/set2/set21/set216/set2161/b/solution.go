package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		ok := drive(reader)
		if ok {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []string) bool {
	n := len(a)
	// n * n
	var arr [][]int
	for i := range n {
		for j := range n {
			if a[i][j] == '#' {
				arr = append(arr, []int{i, j})
			}
		}
	}

	if len(arr) == 0 {
		return true
	}

	if len(arr) == 4 {
		// 这种情况后面的检查无法处理
		r0, c0 := arr[0][0], arr[0][1]
		if r0+1 < n && c0+1 < n && a[r0+1][c0] == '#' && a[r0+1][c0+1] == '#' && a[r0][c0+1] == '#' {
			return true
		}
	}

	var travel func(r int, c int, cnt int, d int, w [][]int) bool

	travel = func(r int, c int, cnt int, d int, w [][]int) bool {
		if r < 0 || r >= n || c < 0 || c >= n {
			return false
		}
		if a[r][c] == '#' {
			cnt++
		}
		if cnt == len(arr) {
			return true
		}
		nr, nc := r+w[d][0], c+w[d][1]
		return travel(nr, nc, cnt, d^1, w)
	}

	check := func(sr int, sc int) bool {
		// 先右后下
		if travel(sr, sc, 0, 0, [][]int{{0, 1}, {1, 0}}) {
			return true
		}
		// 先下后右
		if travel(sr, sc, 0, 0, [][]int{{1, 0}, {0, 1}}) {
			return true
		}
		// 先左后下
		if travel(sr, sc, 0, 0, [][]int{{0, -1}, {1, 0}}) {
			return true
		}
		if travel(sr, sc, 0, 0, [][]int{{1, 0}, {0, -1}}) {
			return true
		}
		return false
	}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			r, c := arr[0][0]+dx, arr[0][1]+dy
			if r >= 0 && r < n && c >= 0 && c < n && check(r, c) {
				return true
			}
		}
	}

	return false
}
