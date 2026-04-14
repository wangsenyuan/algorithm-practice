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
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n+1)
	for i := range n + 1 {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	return solve(points)
}

func solve(points [][]int) int {
	n := len(points)

	x, y := points[0][0], points[0][1]

	var res int
	for i := 1; i+1 < n; i++ {
		x0, y0 := points[i][0], points[i][1]
		x1, y1 := points[i+1][0], points[i+1][1]

		// 水始终在右手边
		if x == x0 {
			// same vertical line
			// x1 != x0, and y1 == y0
			if y < y0 && x1 < x0 {
				res++
			}
			if y > y0 && x1 > x0 {
				res++
			}
		} else {
			if x < x0 && y1 > y0 {
				res++
			}
			if x > x0 && y1 < y0 {
				res++
			}
		}
		x, y = x0, y0
	}

	return res
}
