package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var a, b, x1, y1, x2, y2 int
	fmt.Fscan(reader, &a, &b, &x1, &y1, &x2, &y2)
	return solve(a, b, x1, y1, x2, y2)
}

func solve(a int, b int, x1 int, y1 int, x2 int, y2 int) int {
	u1 := floorDiv(x1+y1, 2*a)
	u2 := floorDiv(x2+y2, 2*a)
	v1 := floorDiv(x1-y1, 2*b)
	v2 := floorDiv(x2-y2, 2*b)

	return max(abs(u1-u2), abs(v1-v2))
}

func floorDiv(x int, d int) int {
	q := x / d
	r := x % d
	if r < 0 {
		q--
	}
	return q
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
