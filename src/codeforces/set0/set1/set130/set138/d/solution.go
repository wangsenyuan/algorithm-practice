package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) string {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	field := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &field[i])
	}
	return solve(field)
}

func solve(field []string) string {
	n := len(field)
	m := len(field[0])
	limit := n + m - 2
	shift := m - 1

	type cell struct {
		r  int
		c  int
		u  int
		v  int
		ch byte
	}

	cells := make([][]cell, 2)
	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			p := (r + c) & 1
			cells[p] = append(cells[p], cell{
				r:  r,
				c:  c,
				u:  r + c,
				v:  r - c + shift,
				ch: field[r][c],
			})
		}
	}

	type solver struct {
		p    int
		memo map[int]int
	}

	var key = func(ul, ur, vl, vr int) int {
		return (((ul*40+ur)*40+vl)*40 + vr)
	}

	var calc func(*solver, int, int, int, int) int

	calc = func(s *solver, ul, ur, vl, vr int) int {
		if ul > ur || vl > vr {
			return 0
		}
		k := key(ul, ur, vl, vr)
		if v, ok := s.memo[k]; ok {
			return v
		}

		seen := make([]bool, 128)
		any := false

		addSeen := func(x int) {
			if x >= len(seen) {
				tmp := make([]bool, x+1)
				copy(tmp, seen)
				seen = tmp
			}
			seen[x] = true
		}

		for _, cur := range cells[s.p] {
			if cur.u < ul || cur.u > ur || cur.v < vl || cur.v > vr {
				continue
			}
			any = true
			var g int
			if cur.ch == 'L' {
				g = calc(s, ul, cur.u-1, vl, vr) ^ calc(s, cur.u+1, ur, vl, vr)
			} else if cur.ch == 'R' {
				g = calc(s, ul, ur, vl, cur.v-1) ^ calc(s, ul, ur, cur.v+1, vr)
			} else {
				g = calc(s, ul, cur.u-1, vl, cur.v-1) ^
					calc(s, ul, cur.u-1, cur.v+1, vr) ^
					calc(s, cur.u+1, ur, vl, cur.v-1) ^
					calc(s, cur.u+1, ur, cur.v+1, vr)
			}
			addSeen(g)
		}

		if !any {
			s.memo[k] = 0
			return 0
		}

		var mex int
		for mex < len(seen) && seen[mex] {
			mex++
		}
		s.memo[k] = mex
		return mex
	}

	var xor int
	for p := 0; p < 2; p++ {
		s := &solver{p: p, memo: make(map[int]int)}
		xor ^= calc(s, 0, limit, 0, limit)
	}

	if xor != 0 {
		return "WIN"
	}
	return "LOSE"
}
