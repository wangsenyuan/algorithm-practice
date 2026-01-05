package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	best, row, col := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, best)
	if best < 0 {
		return
	}
	for i, v := range row {
		for range v {
			fmt.Fprintln(writer, "row ", i+1)
		}
	}
	for i, v := range col {
		for range v {
			fmt.Fprintln(writer, "col ", i+1)
		}
	}
}

func drive(reader *bufio.Reader) (best int, row []int, col []int) {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, m)
		for j := range m {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	best, row, col = solve(a)
	return
}

func solve(a [][]int) (int, []int, []int) {
	n := len(a)
	m := len(a[0])

	x0 := slices.Min(a[0])
	row := make([]int, n)
	col := make([]int, m)

	play := func() int {
		for j := range m {
			col[j] = a[0][j] - row[0]
			if col[j] < 0 {
				return -1
			}
		}
		for i := 1; i < n; i++ {
			row[i] = a[i][0] - col[0]
			if row[i] < 0 {
				return -1
			}
		}

		for i := range n {
			for j := range m {
				if a[i][j] != row[i]+col[j] {
					return -1
				}
			}
		}

		var sum int
		for i := range n {
			sum += row[i]
		}
		for i := range m {
			sum += col[i]
		}

		return sum
	}

	best := -1
	var x2 int

	for x := range x0 + 1 {
		row[0] = x
		tmp := play()
		if tmp < 0 {
			continue
		}
		if best < 0 || tmp < best {
			best = tmp
			x2 = x
		}
	}

	row[0] = x2
	play()

	return best, row, col
}
