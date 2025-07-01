package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	k, maze, res := process(reader)

	var cnt int

	var sum int
	for i := range maze {
		for j := range maze[i] {
			if maze[i][j] != res[i][j] {
				cnt++
			}
			if res[i][j] == '.' {
				sum++
			}
		}
	}
	if cnt != k {
		t.Fatalf("Sample result %v, not correct, expect change %d cells, but got %d", res, k, cnt)
	}

	marked := make([][]bool, len(maze))
	for i := range maze {
		marked[i] = make([]bool, len(maze[i]))
	}

	var dfs func(r int, c int)
	dfs = func(r int, c int) {
		sum--
		marked[r][c] = true
		for i := range 4 {
			x, y := r+dd[i], c+dd[i+1]
			if x >= 0 && x < len(res) && y >= 0 && y < len(res[x]) && res[x][y] == '.' && !marked[x][y] {
				dfs(x, y)
			}
		}
	}

	for i := range res {
		for j := range res[i] {
			if res[i][j] == '.' && sum > 0 {
				dfs(i, j)
				if sum != 0 {
					t.Fatalf("Sample result %v, not correct", res)
				}
			}
		}
	}

}

func TestSample1(t *testing.T) {
	runSample(t, `3 4 2
#..#
..#.
#...
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 4 5
#...
#.#.
.#..
...#
.#.#
`)
}
