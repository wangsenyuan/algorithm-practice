package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int64) {
	reader := bufio.NewReader(strings.NewReader(s))
	cnt, field := drive(reader)
	if cnt != expect {
		t.Fatalf("Sample expect count %d, but got %d", expect, cnt)
	}
	checkAnswer(t, strings.Split(strings.TrimSpace(s), "\n")[1:], field)
}

func checkAnswer(t *testing.T, input []string, out []string) {
	t.Helper()
	n := len(input)
	m := len(input[0])
	for i := range n {
		if len(out[i]) != m {
			t.Fatalf("bad output row length: %q", out[i])
		}
		for j := range m {
			if input[i][j] == '.' {
				if out[i][j] != '.' {
					t.Fatalf("empty cell (%d,%d) changed to %c", i, j, out[i][j])
				}
			} else if out[i][j] < '0' || out[i][j] > '6' {
				t.Fatalf("occupied cell (%d,%d) has bad digit %c", i, j, out[i][j])
			}
		}
	}

	used := make(map[[2]byte]bool)
	pos := make(map[byte][]cell)
	for i := range n {
		for j := range m {
			if input[i][j] != '.' {
				pos[input[i][j]] = append(pos[input[i][j]], cell{i, j})
			}
		}
	}
	for _, ps := range pos {
		a := out[ps[0].r][ps[0].c]
		b := out[ps[1].r][ps[1].c]
		if a > b {
			a, b = b, a
		}
		key := [2]byte{a, b}
		if used[key] {
			t.Fatalf("duplicate domino %c-%c", a, b)
		}
		used[key] = true
	}
	if len(used) != 28 {
		t.Fatalf("expected 28 different dominoes, got %d", len(used))
	}

	covered := make([][]bool, n)
	for i := range n {
		covered[i] = make([]bool, m)
	}
	for {
		r, c := -1, -1
		for i := 0; i < n && r < 0; i++ {
			for j := 0; j < m; j++ {
				if input[i][j] != '.' && !covered[i][j] {
					r, c = i, j
					break
				}
			}
		}
		if r < 0 {
			break
		}
		d := out[r][c]
		for x := r; x < r+2; x++ {
			for y := c; y < c+2; y++ {
				if x >= n || y >= m || input[x][y] == '.' || covered[x][y] || out[x][y] != d {
					t.Fatalf("bad 2x2 square at (%d,%d)", r, c)
				}
				covered[x][y] = true
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `8 8
.aabbcc.
.defghi.
kdefghij
klmnopqj
.lmnopq.
.rstuvw.
xrstuvwy
xzzAABBy`
	runSample(t, s, 10080)
}
