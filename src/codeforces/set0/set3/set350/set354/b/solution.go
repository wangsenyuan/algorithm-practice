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

	fmt.Fprintln(writer, drive(reader))
}

func drive(reader *bufio.Reader) string {
	var n int
	fmt.Fscan(reader, &n)

	grid := make([]string, n)
	for i := range grid {
		fmt.Fscan(reader, &grid[i])
	}

	return solve(grid)
}

func score(c byte) int {
	if c == 'a' {
		return 1
	}
	if c == 'b' {
		return -1
	}
	return 0
}

const inf = 1 << 60

func solve(grid []string) string {
	n := len(grid)
	diagLen := func(d int) int {
		if d < n {
			return d + 1
		}
		return 2*n - 1 - d
	}

	// On the terminal diagonal, the only non-empty state is the bottom-right cell.
	next := []int{0, 0}

	for d := 2*n - 3; d >= 0; d-- {
		length := diagLen(d)
		cur := make([]int, 1<<length)
		firstRow := max(0, d-n+1)

		for state := 1; state < len(cur); state++ {
			var byLetter [26]int
			for k := 0; k < length; k++ {
				if state&(1<<k) == 0 {
					continue
				}
				r := firstRow + k
				c := d - r

				if c+1 < n {
					pos := k
					if d >= n-1 {
						pos--
					}
					byLetter[grid[r][c+1]-'a'] |= 1 << pos
				}
				if r+1 < n {
					pos := k + 1
					if d >= n-1 {
						pos--
					}
					byLetter[grid[r+1][c]-'a'] |= 1 << pos
				}
			}

			best := -inf
			if d&1 == 0 {
				best = inf
			}
			for ch, nextState := range byLetter {
				if nextState == 0 {
					continue
				}
				candidate := score(byte('a'+ch)) + next[nextState]
				if d&1 == 0 {
					best = min(best, candidate)
				} else {
					best = max(best, candidate)
				}
			}
			cur[state] = best
		}

		next = cur
	}

	res := score(grid[0][0]) + next[1]
	if res > 0 {
		return "FIRST"
	}
	if res < 0 {
		return "SECOND"
	}
	return "DRAW"
}
