package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%d %d %d %d\n", res[0], res[1], res[2], res[3])
}

func drive(reader *bufio.Reader) []int {
	var n, m, x, y, a, b int
	fmt.Fscan(reader, &n, &m, &x, &y, &a, &b)
	return solve(n, m, x, y, a, b)
}

func solve(n int, m int, x int, y int, a int, b int) []int {
	g := gcd(a, b)
	a /= g
	b /= g

	scale := min(n/a, m/b)
	w := a * scale
	h := b * scale

	x1 := place(n, x, w)
	y1 := place(m, y, h)

	return []int{x1, y1, x1 + w, y1 + h}
}

func place(limit int, pos int, length int) int {
	low := max(0, pos-length)
	high := min(pos, limit-length)

	cand := (2*pos - length) / 2
	if cand < low {
		return low
	}
	if cand > high {
		return high
	}
	return cand
}

func gcd(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
