package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	points := make([][]int, 3)
	for i := range 3 {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		points[i] = []int{x, y}
	}
	return solve(points)
}

func solve(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	if points[0][0] == points[1][0] && points[0][0] == points[2][0] {
		return 1
	}
	if points[0][1] == points[1][1] && points[0][1] == points[2][1] {
		return 1
	}

	if points[0][0] == points[1][0] && (points[2][1] <= points[0][1] || points[2][1] >= points[1][1]) {
		return 2
	}

	// 这个时候第三个节点肯定在第二个点的后面
	if points[0][1] == points[1][1] || points[1][1] == points[2][1] {
		return 2
	}

	if points[1][0] == points[2][0] && (points[0][1] <= points[1][1] || points[0][1] >= points[2][1]) {
		return 2
	}

	return 3
}
