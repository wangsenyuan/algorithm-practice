package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%.12f\n", drive(reader))
}

func drive(reader *bufio.Reader) float64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	universe := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &universe[i])
	}
	return solve(universe)
}

func solve(universe []string) float64 {
	n := len(universe)
	m := len(universe[0])

	row := make([]int64, n)
	col := make([]int64, m)
	var free int64

	for i := range n {
		for j := range m {
			if universe[i][j] != 'X' {
				row[i]++
				col[j]++
				free++
			}
		}
	}

	var tot int64

	var cnt, sum int64
	for i := range n {
		tot += int64(i)*cnt*row[i] - sum*row[i]
		cnt += row[i]
		sum += int64(i) * row[i]
	}

	cnt, sum = 0, 0
	for j := range m {
		tot += int64(j)*cnt*col[j] - sum*col[j]
		cnt += col[j]
		sum += int64(j) * col[j]
	}

	for i := range n {
		var left int64
		for j := range m {
			if universe[i][j] == 'X' {
				tot += 2 * left * (row[i] - left)
			} else {
				left++
			}
		}
	}

	for j := range m {
		var up int64
		for i := range n {
			if universe[i][j] == 'X' {
				tot += 2 * up * (col[j] - up)
			} else {
				up++
			}
		}
	}

	return float64(tot*2) / float64(free*free)
}
