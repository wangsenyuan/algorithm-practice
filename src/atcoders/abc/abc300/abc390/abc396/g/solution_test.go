package main

import (
	"bufio"
	"math/bits"
	"math/rand"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3
100
010
110
`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4
1111
1111
1111
`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `10 5
10000
00111
11000
01000
10110
01110
10101
00100
00100
10001
`, 13)
}

func bruteForce(w int, grid []string) int {
	rows := make([]int, len(grid))
	for i, row := range grid {
		for _, c := range row {
			rows[i] = rows[i]<<1 | int(c-'0')
		}
	}
	best := len(grid) * w
	for mask := 0; mask < 1<<w; mask++ {
		cur := 0
		for _, row := range rows {
			ones := bits.OnesCount(uint(row ^ mask))
			cur += min(ones, w-ones)
		}
		best = min(best, cur)
	}
	return best
}

func TestRandomizedAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(396))
	for tc := 0; tc < 1000; tc++ {
		h := rng.Intn(7) + 1
		w := rng.Intn(7) + 1
		grid := make([]string, h)
		for i := range grid {
			row := make([]byte, w)
			for j := range row {
				row[j] = byte('0' + rng.Intn(2))
			}
			grid[i] = string(row)
		}
		got := solve(w, grid)
		want := bruteForce(w, grid)
		if got != want {
			t.Fatalf("w=%d grid=%v: got %d, want %d", w, grid, got, want)
		}
	}
}
