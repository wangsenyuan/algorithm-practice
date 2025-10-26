package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res < 0 {
		fmt.Println("IMPOSSIBLE")
		return
	}
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	notes := make([][]int, m)
	for i := 0; i < m; i++ {
		notes[i] = make([]int, 2)
		fmt.Fscan(reader, &notes[i][0], &notes[i][1])
	}
	return solve(n, m, notes)
}

func solve(n int, m int, notes [][]int) int {

	calc := func(prev []int, cur []int) int {
		day_diff := cur[0] - prev[0]
		if cur[1] > prev[1] && day_diff < cur[1]-prev[1] {
			return -1
		}
		if cur[1] < prev[1] && day_diff < prev[1]-cur[1] {
			return -1
		}
		// 假设一直增长到x, x - prev[1] + x - cur[1] = day_diff
		// 0 1 2 3 2 1
		return (prev[1] + cur[1] + day_diff) / 2
	}
	res := notes[0][1] + notes[0][0] - 1

	for i := 0; i+1 < m; i++ {
		tmp := calc(notes[i], notes[i+1])
		if tmp < 0 {
			return -1
		}
		res = max(res, tmp, notes[i][1])
	}

	res = max(res, notes[m-1][1]+n-notes[m-1][0])

	return res
}
