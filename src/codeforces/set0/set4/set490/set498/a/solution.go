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
	var x1, y1, x2, y2 int
	fmt.Fscan(reader, &x1, &y1, &x2, &y2)
	var n int
	fmt.Fscan(reader, &n)
	roads := make([][]int, n)
	for i := range n {
		roads[i] = make([]int, 3)
		fmt.Fscan(reader, &roads[i][0], &roads[i][1], &roads[i][2])
	}
	return solve([]int{x1, y1}, []int{x2, y2}, roads)
}

func solve(first []int, second []int, roads [][]int) int {
	// a * x + b * y = c
	var res int

	for _, road := range roads {
		a, b, c := road[0], road[1], road[2]
		// Line a*x + b*y + c = 0: points on opposite sides iff sign(a*x+b*y+c) differs
		v1 := a*first[0] + b*first[1] + c
		v2 := a*second[0] + b*second[1] + c
		if sign(v1) != sign(v2) {
			res++
		}
	}
	return res
}

func sign(num int) int {
	if num > 0 {
		return 1
	}
	if num < 0 {
		return -1
	}
	return 0
}
