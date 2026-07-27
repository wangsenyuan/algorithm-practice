package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []string) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res, k := drive(reader)

	if len(expect) > 0 != (len(res) > 0) {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if len(expect) == 0 {
		return
	}

	if len(res) != 2 {
		t.Fatalf("Sample expect 2 rows, but got %v", res)
	}
	if len(expect[0]) != len(res[0]) || len(res[0]) != len(res[1]) {
		t.Fatalf("Sample expect width %d, but got %v", len(expect[0]), res)
	}

	countMineNeighbors := func(i, j int) int {
		var cnt int
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				r, c := i+dx, j+dy
				if r >= 0 && r < 2 && c >= 0 && c < len(res[0]) && res[r][c] == '*' {
					cnt++
				}
			}
		}
		return cnt
	}

	var cnt int
	for i := range 2 {
		for j := range len(res[0]) {
			if res[i][j] != '.' {
				continue
			}
			mines := countMineNeighbors(i, j)
			if mines > 1 {
				t.Fatalf("Sample result %v is invalid: empty (%d,%d) has %d mine neighbors", res, i, j, mines)
			}
			if mines >= 1 {
				cnt++
			}
		}
	}
	if cnt != k {
		t.Fatalf("Sample expect %d threatened empties, but got %d; grid %v", k, cnt, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `1
`, []string{"*", "."})
}

func TestSample2(t *testing.T) {
	runSample(t, `4
`, nil)
}

func TestSample3(t *testing.T) {
	runSample(t, `8
`, []string{"*....", "...*."})
}

func TestSample4(t *testing.T) {
	runSample(t, `10
`, []string{".*..*.", "......"})
}

func TestSample5(t *testing.T) {
	runSample(t, `9
`, nil)
}
