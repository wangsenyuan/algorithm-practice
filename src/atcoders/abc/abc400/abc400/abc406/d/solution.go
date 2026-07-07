package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for _, ans := range drive(reader) {
		fmt.Println(ans)
	}
}

func drive(reader *bufio.Reader) []int {
	var h, w, n int
	fmt.Fscan(reader, &h, &w, &n)
	trash := make([][]int, n)
	for i := range n {
		trash[i] = make([]int, 2)
		fmt.Fscan(reader, &trash[i][0], &trash[i][1])
	}
	var q int
	fmt.Fscan(reader, &q)
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(h, w, trash, queries)
}

func solve(h, w int, trash [][]int, queries [][]int) []int {
	rowTodo := make([]map[int]bool, h+1)
	for i := range h + 1 {
		rowTodo[i] = make(map[int]bool)
	}
	colTodo := make([]map[int]bool, w+1)
	for j := range w + 1 {
		colTodo[j] = make(map[int]bool)
	}

	for i, cur := range trash {
		r, c := cur[0], cur[1]
		rowTodo[r][i] = true
		colTodo[c][i] = true
	}

	ans := make([]int, len(queries))
	for i, cur := range queries {
		if cur[0] == 1 {
			r := cur[1]
			ans[i] = len(rowTodo[r])
			for j := range rowTodo[r] {
				c := trash[j][1]
				delete(colTodo[c], j)
			}
			clear(rowTodo[r])
		} else {
			c := cur[1]
			ans[i] = len(colTodo[c])
			for j := range colTodo[c] {
				r := trash[j][0]
				delete(rowTodo[r], j)
			}
			clear(colTodo[c])
		}
	}

	return ans
}
