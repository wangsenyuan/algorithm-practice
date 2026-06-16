package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for _, poly := range res {
		fmt.Fprintln(writer, len(poly))
		for _, p := range poly {
			fmt.Fprintln(writer, p.x, p.y)
		}
	}
}

func drive(reader *bufio.Reader) [][]Point {
	var res [][]Point
	for {
		var n int
		if _, err := fmt.Fscan(reader, &n); err != nil || n == 0 {
			break
		}
		grid := make([][]int, n+1)
		for i := range grid {
			grid[i] = make([]int, n+1)
		}
		for y := n; y >= 1; y-- {
			var s string
			fmt.Fscan(reader, &s)
			for x := 1; x <= n; x++ {
				grid[x][y] = int(s[x-1] - '0')
			}
		}
		res = append(res, solve(grid))
	}
	return res
}

func solve(grid [][]int) []Point {
	n := len(grid) - 1
	inside := make([][]int, n+2)
	for i := range inside {
		inside[i] = make([]int, n+2)
	}

	for y := 1; y <= n; y++ {
		for x := 1; x <= n; x++ {
			inside[x+1][y+1] = grid[x][y] - inside[x][y] - inside[x+1][y] - inside[x][y+1]
		}
	}

	var points []Point
	for x := 1; x <= n+1; x++ {
		for y := 1; y <= n+1; y++ {
			if inside[x][y] == 1 {
				points = append(points, Point{x - 1, y - 1})
			}
		}
	}

	return clockwise(hull(points))
}

func hull(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})

	if len(points) <= 1 {
		return points
	}

	var lower []Point
	for _, p := range points {
		for len(lower) >= 2 && cross(lower[len(lower)-2], lower[len(lower)-1], p) <= 0 {
			lower = lower[:len(lower)-1]
		}
		lower = append(lower, p)
	}

	var upper []Point
	for i := len(points) - 1; i >= 0; i-- {
		p := points[i]
		for len(upper) >= 2 && cross(upper[len(upper)-2], upper[len(upper)-1], p) <= 0 {
			upper = upper[:len(upper)-1]
		}
		upper = append(upper, p)
	}

	return append(lower[:len(lower)-1], upper[:len(upper)-1]...)
}

func clockwise(poly []Point) []Point {
	if len(poly) <= 2 {
		return poly
	}
	res := make([]Point, 0, len(poly))
	res = append(res, poly[0])
	for i := len(poly) - 1; i >= 1; i-- {
		res = append(res, poly[i])
	}
	return res
}

func cross(a Point, b Point, c Point) int {
	x1, y1 := b.x-a.x, b.y-a.y
	x2, y2 := c.x-a.x, c.y-a.y
	return x1*y2 - y1*x2
}
