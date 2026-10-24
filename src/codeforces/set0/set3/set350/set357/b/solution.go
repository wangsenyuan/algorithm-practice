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

	s := fmt.Sprintf("%v", drive(reader))
	fmt.Fprintln(writer, s[1:len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	dances := make([][]int, m)
	for i := range m {
		dances[i] = make([]int, 3)
		fmt.Fscan(reader, &dances[i][0], &dances[i][1], &dances[i][2])
	}
	return solve(n, dances)
}

func solve(n int, dances [][]int) []int {
	assign := make([]int, n)

	for _, cur := range dances {
		i, j, k := cur[0]-1, cur[1]-1, cur[2]-1
		if assign[i]+assign[j]+assign[k] == 0 {
			assign[i] = 1
			assign[j] = 2
			assign[k] = 3
		} else {
			// 只有一个不为0
			if assign[j] != 0 {
				i, j = j, i
			} else if assign[k] != 0 {
				i, k = k, i
			}
			// assign[i] > 0
			if assign[i] < 3 {
				assign[j] = assign[i] + 1
			} else {
				assign[j] = 1
			}
			assign[k] = 6 - assign[i] - assign[j]
		}
	}

	return assign
}
