package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var byRow []point
	var byCol []point
	var free int64

	for i := range n {
		for j := range m {
			if universe[i][j] != 'X' {
				row[i]++
				col[j]++
				free++
			} else {
				byRow = append(byRow, point{i, j})
				byCol = append(byCol, point{j, i})
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

	sort.Slice(byCol, func(i, j int) bool {
		return byCol[i].x < byCol[j].x
	})
	tot += countChains(byRow, m)
	tot += countChains(byCol, n)

	return float64(tot*2) / float64(free*free)
}

type point struct {
	x int
	y int
}

func countChains(a []point, limit int) int64 {
	var res int64

	for l := 0; l < len(a); {
		r := l + 1
		for r < len(a) && a[r].x == a[r-1].x+1 {
			r++
		}

		for i := l; i < r; i++ {
			dir := 0
			for j := i + 1; j < r; j++ {
				cur := 1
				if a[j].y < a[j-1].y {
					cur = -1
				}
				if dir == 0 {
					dir = cur
				} else if dir != cur {
					break
				}
				if dir > 0 {
					res += 2 * int64(a[i].y) * int64(limit-1-a[j].y)
				} else {
					res += 2 * int64(limit-1-a[i].y) * int64(a[j].y)
				}
			}
		}

		l = r
	}

	return res
}
