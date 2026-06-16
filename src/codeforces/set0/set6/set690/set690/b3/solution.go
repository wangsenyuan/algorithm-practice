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
		var n, m int
		if _, err := fmt.Fscan(reader, &n, &m); err != nil || n == 0 && m == 0 {
			break
		}
		cells := make([]Point, m)
		for i := range cells {
			fmt.Fscan(reader, &cells[i].x, &cells[i].y)
		}
		res = append(res, solve(n, cells))
	}
	return res
}

func solve(n int, cells []Point) []Point {
	known := make(map[Point]bool, len(cells)*2)
	centers := make([]Point, len(cells))
	for i, p := range cells {
		known[p] = true
		centers[i] = Point{2*p.x - 1, 2*p.y - 1}
	}
	centerHull := hull(centers)

	isNonZero := func(c Point) bool {
		if c.x < 1 || c.x > n || c.y < 1 || c.y > n {
			return false
		}
		if known[c] {
			return true
		}
		return insideConvex(centerHull, Point{2*c.x - 1, 2*c.y - 1})
	}

	candidates := make(map[Point]bool)
	for _, c := range cells {
		corners := []Point{
			{c.x - 1, c.y - 1},
			{c.x, c.y - 1},
			{c.x - 1, c.y},
			{c.x, c.y},
		}
		for _, p := range corners {
			if p.x < 1 || p.x >= n || p.y < 1 || p.y >= n {
				continue
			}
			ok := true
			for dx := 0; dx <= 1 && ok; dx++ {
				for dy := 0; dy <= 1; dy++ {
					if !isNonZero(Point{p.x + dx, p.y + dy}) {
						ok = false
						break
					}
				}
			}
			if ok {
				candidates[p] = true
			}
		}
	}

	points := make([]Point, 0, len(candidates))
	for p := range candidates {
		points = append(points, p)
	}
	return clockwise(hull(points))
}

func insideConvex(poly []Point, p Point) bool {
	n := len(poly)
	if n == 0 {
		return false
	}
	if n == 1 {
		return poly[0] == p
	}
	if n == 2 {
		return cross(poly[0], poly[1], p) == 0 &&
			min(poly[0].x, poly[1].x) <= p.x && p.x <= max(poly[0].x, poly[1].x) &&
			min(poly[0].y, poly[1].y) <= p.y && p.y <= max(poly[0].y, poly[1].y)
	}
	if cross(poly[0], poly[1], p) < 0 {
		return false
	}
	if cross(poly[0], poly[n-1], p) > 0 {
		return false
	}
	l, r := 1, n-1
	for r-l > 1 {
		mid := (l + r) / 2
		if cross(poly[0], poly[mid], p) >= 0 {
			l = mid
		} else {
			r = mid
		}
	}
	return cross(poly[l], poly[l+1], p) >= 0
}

func hull(points []Point) []Point {
	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		}
		return points[i].y < points[j].y
	})
	points = unique(points)

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

func unique(points []Point) []Point {
	res := points[:0]
	for _, p := range points {
		if len(res) == 0 || res[len(res)-1] != p {
			res = append(res, p)
		}
	}
	return res
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

func cross(a Point, b Point, c Point) int64 {
	x1, y1 := int64(b.x-a.x), int64(b.y-a.y)
	x2, y2 := int64(c.x-a.x), int64(c.y-a.y)
	return x1*y2 - y1*x2
}
