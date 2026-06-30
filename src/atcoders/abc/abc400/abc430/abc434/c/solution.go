package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		if drive(reader) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	var h int
	fmt.Fscan(reader, &n, &h)
	goals := make([][]int, n)
	for i := range n {
		goals[i] = make([]int, 3)
		fmt.Fscan(reader, &goals[i][0], &goals[i][1], &goals[i][2])
	}
	return solve(h, goals)
}

func solve(h int, goals [][]int) bool {
	lo, hi := h, h

	var prev int
	for _, goal := range goals {
		t, l, r := goal[0], goal[1], goal[2]
		l1 := max(1, lo-(t-prev))
		r1 := hi + t - prev
		if r < l1 || r1 < l {
			return false
		}
		lo = max(l1, l)
		hi = min(r1, r)
		prev = t
	}

	return true
}
