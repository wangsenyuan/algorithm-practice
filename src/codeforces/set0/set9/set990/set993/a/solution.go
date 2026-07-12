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

func drive(reader *bufio.Reader) string {
	axis := readSquare(reader)
	rotated := readSquare(reader)
	return solve(axis, rotated)
}

func readSquare(reader *bufio.Reader) [][]int {
	sq := make([][]int, 4)
	for i := range 4 {
		sq[i] = make([]int, 2)
		fmt.Fscan(reader, &sq[i][0], &sq[i][1])
	}
	return sq
}

func solve(axis [][]int, rotated [][]int) string {
	if check(axis, rotated) {
		return "YES"
	}
	if check(rotate(rotated), rotate(axis)) {
		return "YES"
	}
	return "NO"
}

func rotate(square [][]int) [][]int {
	res := make([][]int, 4)
	for i, p := range square {
		res[i] = []int{p[0] + p[1], p[0] - p[1]}
	}
	return res
}

func check(axis [][]int, rotated [][]int) bool {
	minX, maxX := axis[0][0], axis[0][0]
	minY, maxY := axis[0][1], axis[0][1]
	for _, p := range axis {
		minX = min(minX, p[0])
		maxX = max(maxX, p[0])
		minY = min(minY, p[1])
		maxY = max(maxY, p[1])
	}

	checkInside := func(point []int) bool {
		return minX <= point[0] && point[0] <= maxX && minY <= point[1] && point[1] <= maxY
	}

	// 检查rotated的4个的顶点, 是否在 axis表示的内部(边)
	for _, corner := range rotated {
		if checkInside(corner) {
			return true
		}
	}

	var sx, sy int
	for _, p := range rotated {
		sx += p[0]
		sy += p[1]
	}
	if 4*minX <= sx && sx <= 4*maxX && 4*minY <= sy && sy <= 4*maxY {
		return true
	}

	return false
}
