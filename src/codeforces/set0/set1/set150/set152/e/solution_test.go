package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expectCost int, _ []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	cost, grid, a, important := drive(reader)
	if cost != expectCost {
		t.Fatalf("Sample expect cost=%v, but got cost=%v grid=%v",
			expectCost, cost, grid)
	}

	validatePlan(t, cost, grid, a, important)
}

func validatePlan(t *testing.T, cost int, grid []string, a [][]int, important [][]int) {
	t.Helper()
	n := len(a)
	m := len(a[0])
	if len(grid) != n {
		t.Fatalf("result has %d rows, want %d", len(grid), n)
	}

	var sum int
	for i := range n {
		if len(grid[i]) != m {
			t.Fatalf("row %d has length %d, want %d", i, len(grid[i]), m)
		}
		for j := range m {
			switch grid[i][j] {
			case 'X':
				sum += a[i][j]
			case '.':
			default:
				t.Fatalf("unexpected character %q at (%d,%d)", grid[i][j], i, j)
			}
		}
	}
	if sum != cost {
		t.Fatalf("result grid costs %d, but solver returned %d", sum, cost)
	}

	for _, cur := range important {
		r, c := cur[0]-1, cur[1]-1
		if grid[r][c] != 'X' {
			t.Fatalf("important cell (%d,%d) is not covered", r, c)
		}
	}

	startR, startC := important[0][0]-1, important[0][1]-1
	seen := make([][]bool, n)
	for i := range n {
		seen[i] = make([]bool, m)
	}
	queue := [][2]int{{startR, startC}}
	seen[startR][startC] = true
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for head := 0; head < len(queue); head++ {
		r, c := queue[head][0], queue[head][1]
		for _, d := range dirs {
			x, y := r+d[0], c+d[1]
			if x >= 0 && x < n && y >= 0 && y < m &&
				!seen[x][y] && grid[x][y] == 'X' {
				seen[x][y] = true
				queue = append(queue, [2]int{x, y})
			}
		}
	}
	for _, cur := range important {
		r, c := cur[0]-1, cur[1]-1
		if !seen[r][c] {
			t.Fatalf("important cell (%d,%d) is disconnected in %v", r, c, grid)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 2
1 2 3
1 2 3
1 2 3
1 2
3 3
`, 9, []string{
		".X.",
		".X.",
		".XX",
	})
}

func TestSample2(t *testing.T) {
	runSample(t, `4 5 4
1 4 5 1 2
2 2 2 2 7
2 4 1 4 5
3 2 1 7 1
1 1
1 5
4 1
4 4
`, 26, []string{
		"X..XX",
		"XXXX.",
		"X.X..",
		"X.XX.",
	})
}

func TestSingleTerminal(t *testing.T) {
	a := [][]int{{7}}
	important := [][]int{{1, 1}}
	cost, grid := solve(a, important)
	if cost != 7 {
		t.Fatalf("got cost %d, want 7", cost)
	}
	validatePlan(t, cost, grid, a, important)
}

func TestIntermediatePathCellsAreReconstructed(t *testing.T) {
	a := [][]int{{2, 3, 5}}
	important := [][]int{{1, 1}, {1, 3}}
	cost, grid := solve(a, important)
	if cost != 10 {
		t.Fatalf("got cost %d, want 10", cost)
	}
	validatePlan(t, cost, grid, a, important)
}
