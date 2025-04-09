package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(process(reader))
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) string {
	a := make([]string, 8)
	for i := 0; i < 8; i++ {
		a[i] = readString(reader)
	}
	return solve(a)
}

func solve(a []string) string {
	// 从上到下
	get := func(state uint64, r int, c int) int {
		if r < 0 || r > 7 || c < 0 || c > 7 {
			return 0
		}
		return int((state >> (r*8 + c) & 1))
	}

	var dfs func(state uint64, r int, c int) bool
	dfs = func(state uint64, r int, c int) bool {
		if state == 0 || r == 7 && c == 0 {
			// 接下来随意运行
			return true
		}

		next := state >> 8

		for _, dr := range []int{-1, 0, 1} {
			for _, dc := range []int{-1, 0, 1} {
				nr, nc := r+dr, c+dc
				if nr >= 0 && nr < 8 && nc >= 0 && nc < 8 {
					// nr在当前是安全的，在下一个状态也是安全的
					if get(state, nr, nc) == 0 && get(next, nr, nc) == 0 && dfs(next, nr, nc) {
						return true
					}
				}
			}
		}

		return false
	}

	var state uint64

	for i := 0; i < 8; i++ {
		row := a[7-i]
		for j := 7; j >= 0; j-- {
			if row[j] == 'S' {
				state |= 1 << (i*8 + 7 - j)
			}
		}
	}

	if dfs(state, 0, 7) {
		return "WIN"
	}

	return "LOSE"
}
