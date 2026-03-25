package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, r int
	fmt.Fscan(reader, &n, &r)
	mistakes, ops := solve(n, r)
	if mistakes < 0 {
		fmt.Println("IMPOSSIBLE")
		return
	}
	fmt.Println(mistakes)
	fmt.Println(ops)
}

func solve(n int, r int) (int, string) {
	best := inf
	var ans string

	if n == 1 {
		if r == 1 {
			return 0, "T"
		}
		return -1, ""
	}

	for x := 1; x <= r; x++ {
		if x == r {
			continue
		}
		if steps, mistakes, ok := analyze(r, x, n); ok && mistakes < best {
			best = mistakes
			ans = build(steps)
		}
		if steps, mistakes, ok := analyze(x, r, n); ok && mistakes < best {
			best = mistakes
			ans = build(steps)
		}
	}

	if best == inf {
		return -1, ""
	}
	return best, ans
}

type run struct {
	op  byte
	cnt int
}

func analyze(x int, y int, want int) ([]run, int, bool) {
	var rev []run
	total := 1

	for x != 1 || y != 1 {
		if x <= 0 || y <= 0 {
			return nil, 0, false
		}
		if total > want {
			return nil, 0, false
		}
		if x > y {
			t := (x - 1) / y
			if t == 0 {
				return nil, 0, false
			}
			x -= t * y
			total += t
			rev = append(rev, run{'T', t})
		} else if y > x {
			t := (y - 1) / x
			if t == 0 {
				return nil, 0, false
			}
			y -= t * x
			total += t
			rev = append(rev, run{'B', t})
		} else {
			return nil, 0, false
		}
	}

	if total != want {
		return nil, 0, false
	}

	runs := len(rev)
	if len(rev) == 0 || rev[len(rev)-1].op != 'T' {
		runs++
	}
	mistakes := want - runs
	return rev, mistakes, true
}

func build(rev []run) string {
	var buf strings.Builder
	buf.WriteByte('T')
	for i := len(rev) - 1; i >= 0; i-- {
		for c := 0; c < rev[i].cnt; c++ {
			buf.WriteByte(rev[i].op)
		}
	}
	return buf.String()
}

const inf = 1 << 60
