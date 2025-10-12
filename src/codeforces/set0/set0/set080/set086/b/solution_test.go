package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	a, res := drive(bufio.NewReader(strings.NewReader(s)))
	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
	if !expect {
		return
	}

	n := len(a)
	m := len(a[0])
	marked := make([][]bool, n)
	for i := range n {
		marked[i] = make([]bool, m)
	}

	que := make([]int, n*m)
	// 每个连通块的sz <= 5
	bfs := func(x int, y int) int {
		var head, tail int
		que[head] = x*m + y
		head++
		marked[x][y] = true
		for tail < head {
			r, c := que[tail]/m, que[tail]%m
			tail++
			for i := range 4 {
				nr, nc := r+dd[i], c+dd[i+1]
				if nr >= 0 && nr < n && nc >= 0 && nc < m {
					if a[nr][nc] == '.' {
						if res[nr][nc] == '.' || res[nr][nc] == '#' {
							return -1
						}
						if res[nr][nc] == res[x][y] && !marked[nr][nc] {
							que[head] = nr*m + nc
							head++
							marked[nr][nc] = true
						}
					}
				}
			}
		}
		return head
	}

	for i := range n {
		for j := range m {
			if a[i][j] == '#' || marked[i][j] {
				continue
			}
			sz := bfs(i, j)
			if sz < 2 || sz > 5 {
				t.Fatalf("Sample result %v, causing conflict", res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `2 3
...
#.#`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
.#.
...
..#`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 3
...
.##
.#.`
	expect := false
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `2 3
#..
#.#`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `10 13
##...##..#..#
##......#..##
##......##..#
##..#.....###
#.....##..#..
.###..#...#..
..###.#...##.
#.###.#######
..#..#.######
..###..######`
	expect := true
	runSample(t, s, expect)
}
