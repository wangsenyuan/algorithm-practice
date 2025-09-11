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
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	seals := make([][]int, n)
	for i := 0; i < n; i++ {
		seals[i] = make([]int, 2)
		fmt.Fscan(reader, &seals[i][0], &seals[i][1])
	}
	return solve(x, y, seals)
}

func solve(x int, y int, seals [][]int) int {

	area := func(i int) int {
		return seals[i][0] * seals[i][1]
	}

	check := func(i int, j int) bool {
		if area(i)+area(j) > x*y {
			return false
		}
		a := seals[i]
		b := seals[j]

		for i := range 2 {
			for j := range 2 {
				if a[i]+b[j] <= x && max(a[1^i], b[1^j]) <= y || a[i]+b[j] <= y && max(a[1^i], b[1^j]) <= x {
					return true
				}
			}
		}

		return false
	}

	var best int

	n := len(seals)
	for i := range n {
		for j := range i {
			if check(i, j) {
				best = max(best, area(i)+area(j))
			}
		}
	}
	return best
}
