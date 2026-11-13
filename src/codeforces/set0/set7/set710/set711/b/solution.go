package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)

	grid := make([][]int, n)
	for i := range n {
		grid[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &grid[i][j])
		}
	}

	return solve(grid)
}

func solve(grid [][]int) int {
	if len(grid) == 1 {
		return 1
	}
	var r, c int
	for i, row := range grid {
		for j, v := range row {
			if v == 0 {
				r = i
				c = j
			}
		}
	}

	// 好像会溢出
	sum := big.NewInt(0)
	if r > 0 {
		for _, v := range grid[0] {
			sum = sum.Add(sum, big.NewInt(int64(v)))
		}
	} else {
		for _, v := range grid[1] {
			sum = sum.Add(sum, big.NewInt(int64(v)))
		}
	}

	sum2 := big.NewInt(0)
	for _, v := range grid[r] {
		sum2 = sum2.Add(sum2, big.NewInt(int64(v)))
	}

	w := new(big.Int).Sub(sum, sum2)
	if w.Cmp(big.NewInt(0)) <= 0 || w.Cmp(big.NewInt(1e18)) > 0 {
		return -1
	}
	// check
	col := make([]*big.Int, len(grid))
	for i := range col {
		col[i] = big.NewInt(0)
	}

	for i, row := range grid {
		cur := big.NewInt(0)
		for j, v := range row {
			if i == r && j == c {
				col[j] = col[j].Add(col[j], w)
				cur = cur.Add(cur, w)
			} else {
				col[j] = col[j].Add(col[j], big.NewInt(int64(v)))
				cur = cur.Add(cur, big.NewInt(int64(v)))
			}
		}

		if cur.Cmp(sum) != 0 {
			return -1
		}
	}

	for _, v := range col {
		if v.Cmp(sum) != 0 {
			return -1
		}
	}

	s1, s2 := big.NewInt(0), big.NewInt(0)
	for i := range grid {
		if i == r && i == c {
			s1 = s1.Add(s1, w)
		} else {
			s1 = s1.Add(s1, big.NewInt(int64(grid[i][i])))
		}
		if i == r && len(grid)-1-i == c {
			s2 = s2.Add(s2, w)
		} else {
			s2 = s2.Add(s2, big.NewInt(int64(grid[i][len(grid)-1-i])))
		}
	}

	if s1.Cmp(sum) != 0 || s2.Cmp(sum) != 0 {
		return -1
	}

	return int(w.Int64())
}
