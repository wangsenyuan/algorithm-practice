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

func drive(reader *bufio.Reader) int {
	var s, x1, x2, t1, t2, p, d int
	fmt.Fscan(reader, &s, &x1, &x2, &t1, &t2, &p, &d)
	return solve(s, x1, x2, t1, t2, p, d)
}

func solve(s int, x1 int, x2 int, t1 int, t2 int, p int, d int) int {
	walk := abs(x1-x2) * t2
	if t1 >= t2 {
		return walk
	}

	ride := abs(x1-x2) * t1
	wait := 0

	if x2 > x1 {
		if d == 1 {
			if p <= x1 {
				wait = (x1 - p) * t1
			} else {
				wait = (s - p + s + x1) * t1
			}
		} else {
			wait = (p + x1) * t1
		}
	} else {
		if d == -1 {
			if p >= x1 {
				wait = (p - x1) * t1
			} else {
				wait = (p + s + s - x1) * t1
			}
		} else {
			wait = (s - p + s - x1) * t1
		}
	}

	return min(walk, wait+ride)
}

func abs(num int) int {
	return max(num, -num)
}
