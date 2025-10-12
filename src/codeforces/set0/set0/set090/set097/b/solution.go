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
	points := drive(reader)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(points))
	for _, point := range points {
		fmt.Fprintln(writer, point[0], point[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n int
	fmt.Fscan(reader, &n)

	points := make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}

	return solve(n, points)
}

func solve(n int, points [][]int) [][]int {

	slices.SortFunc(points, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], a[1]-b[1])
	})

	merge := func(a, b [][]int) [][]int {
		var res [][]int
		for i, j := 0, 0; i < len(a) || j < len(b); {
			if j == len(b) || i < len(a) && a[i][1] < b[j][1] {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
		return res
	}

	var f func(pts [][]int) [][]int

	f = func(pts [][]int) [][]int {
		if len(pts) <= 1 {
			return pts
		}

		if len(pts) == 2 {
			res := slices.Clone(pts)
			if res[0][0] != res[1][0] && res[0][1] != res[1][1] {
				res = append(res, []int{pts[1][0], pts[0][1]})
			}
			slices.SortFunc(res, func(a, b []int) int {
				return a[1] - b[1]
			})
			return res
		}

		diff := len(pts)
		mid := -1
		for i := 1; i < len(pts); i++ {
			if pts[i-1][0] < pts[i][0] {
				if abs(len(pts)-i-i) < diff {
					mid = i
					diff = abs(len(pts) - i - i)
				}
			}
		}

		if mid == -1 {
			// 都在同一条垂直线上, 且已经按照y升序排列了
			return pts
		}
		x := pts[mid][0]

		var mid_pts []int
		for i := range pts {
			if pts[i][0] == x {
				mid_pts = append(mid_pts, pts[i][1])
			}
		}

		pref := f(pts[:mid])
		suf := f(pts[mid:])

		res := merge(pref, suf)

		var ys []int
		var j int
		for _, pt := range res {
			for j < len(mid_pts) && mid_pts[j] < pt[1] {
				j++
			}
			if j < len(mid_pts) && mid_pts[j] == pt[1] {
				continue
			}
			if len(ys) == 0 || ys[len(ys)-1] < pt[1] {
				ys = append(ys, pt[1])
			}
		}

		var add [][]int
		for _, y := range ys {
			add = append(add, []int{x, y})
		}

		return merge(res, add)
	}

	return f(points)
}

func abs(num int) int {
	return max(num, -num)
}
