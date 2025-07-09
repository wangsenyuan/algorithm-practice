package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, grid := process(reader)

	m := readNum(reader)
	expect := make([]string, m)
	for i := range m {
		expect[i] = readString(reader)
	}

	buf := make([][]byte, len(grid))

	checkBlueNeighbour := func(i, j int) bool {
		for k := range 4 {
			x, y := i+dd[k], j+dd[k+1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && buf[x][y] == 'B' {
				return true
			}
		}
		return false
	}

	check := func(arr []string) int {
		for i := range len(grid) {
			buf[i] = []byte(grid[i])
		}
		for _, cur := range arr {
			c := cur[0]
			var i, j int
			pos := readInt([]byte(cur), 2, &i) + 1
			readInt([]byte(cur), pos, &j)
			i--
			j--
			switch c {
			case 'B':
				if buf[i][j] != '.' {
					t.Fatalf("Can't build blue tower at (%d, %d)", i+1, j+1)
				}
				buf[i][j] = 'B'
			case 'D':
				if buf[i][j] == '#' || buf[i][j] == '.' {
					t.Fatalf("Can't destroy tower at (%d, %d)", i+1, j+1)
				}
				buf[i][j] = '.'
			default:
				// c == 'R'
				if buf[i][j] != '.' || !checkBlueNeighbour(i, j) {
					t.Fatalf("Can't build red tower at (%d, %d)", i+1, j+1)
				}
				buf[i][j] = 'R'
			}
		}
		var cnt int
		for i := range len(grid) {
			for j := range len(grid[i]) {
				if buf[i][j] == 'B' {
					cnt++
				} else {
					cnt += 2
				}
			}
		}
		return cnt
	}

	a := check(expect)
	b := check(res)

	if a != b {
		t.Errorf("Sample expect %d, but got %d", a, b)
	}

}

func TestSample1(t *testing.T) {
	s := `2 3
..#
.#.
4
B 1 1
R 1 2
R 2 1
B 2 3
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1 3
...
5
B 1 1
B 1 2
R 1 3
D 1 2
R 1 2
`
	runSample(t, s)
}
