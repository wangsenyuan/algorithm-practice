package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expectOK bool) {
	t.Helper()
	grid := parseGrid(s)
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if !expectOK {
		if len(res) != 0 {
			t.Fatalf("Sample expect NO, but got %v", res)
		}
		return
	}
	if len(res) == 0 {
		t.Fatalf("Sample expect YES, but got NO")
	}
	x, y := res[0], res[1]
	if !clearsAll(grid, x, y) {
		t.Fatalf("cell (%d, %d) does not clear every wall", x, y)
	}
}

func parseGrid(s string) []string {
	reader := bufio.NewReader(strings.NewReader(s))
	var n, m int
	fmt.Fscan(reader, &n, &m)
	grid := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &grid[i])
	}
	return grid
}

func clearsAll(grid []string, x, y int) bool {
	n := len(grid)
	if n == 0 || x < 1 || x > n {
		return false
	}
	m := len(grid[0])
	if y < 1 || y > m {
		return false
	}
	for i := range n {
		for j := range m {
			if grid[i][j] == '*' && i != x-1 && j != y-1 {
				return false
			}
		}
	}
	return true
}

func TestSample1(t *testing.T) {
	runSample(t, `3 4
.*..
....
.*..
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 3
..*
.*.
*..
`, false)
}

func TestSample3(t *testing.T) {
	runSample(t, `6 5
..*..
..*..
*****
..*..
..*..
..*..
`, true)
}

func TestSample4(t *testing.T) {
	runSample(t, `1 10
**********
`, true)
}
