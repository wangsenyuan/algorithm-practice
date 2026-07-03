package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	if drive(reader) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n int
	fmt.Fscan(reader, &n)
	a := make([][]int, n)
	for i := range n {
		a[i] = make([]int, 2)
		fmt.Fscan(reader, &a[i][0], &a[i][1])
	}
	var m int
	fmt.Fscan(reader, &m)
	b := make([][]int, m)
	for i := range m {
		b[i] = make([]int, 2)
		fmt.Fscan(reader, &b[i][0], &b[i][1])
	}
	return solve(a, b)
}

const inf = 1 << 60

func solve(a, b [][]int) bool {
	n := len(a)
	m := len(b)
	arr1 := make([]pt, n)
	arr2 := make([]pt, n+m)
	mx := -inf
	for i, cur := range a {
		arr1[i] = pt{cur[0], cur[1]}
		arr2[i] = pt{cur[0], cur[1]}
		mx = max(mx, cur[0])
	}
	for i, cur := range b {
		arr2[n+i] = pt{cur[0], cur[1]}
	}

	first := convex_hull(arr1)
	second := convex_hull(arr2)

	// 这两个应该是相同的
	if len(first) != len(second) {
		return false
	}

	for i, u := range first {
		if u != second[i] {
			return false
		}
	}

	// 找到上不半部分和下半部分
	var upper []pt
	var lower []pt
	for _, cur := range first {
		if len(lower) > 0 {
			lower = append(lower, cur)
		} else {
			upper = append(upper, cur)
			if cur.x == mx {
				lower = append(lower, cur)
			}
		}
	}
	lower = append(lower, first[0])
	slices.Reverse(lower)

	checkOutofLine := func(arr []pt, cur pt) bool {
		r := sort.Search(len(arr), func(i int) bool {
			return arr[i].x > cur.x
		})
		if r == len(arr) || r == 0 {
			// x = mx
			return false
		}

		return orientation(arr[r-1], cur, arr[r]) != 0
	}

	for _, cur := range b {
		pt := pt{cur[0], cur[1]}
		if !checkOutofLine(upper, pt) || !checkOutofLine(lower, pt) {
			return false
		}
	}

	return true
}

type pt struct {
	x int
	y int
}

func orientation(a, b, c pt) int {
	v := int64(a.x)*int64(b.y-c.y) + int64(b.x)*int64(c.y-a.y) + int64(c.x)*int64(a.y-b.y)
	if v < 0 {
		return -1
	}
	if v > 0 {
		return 1
	}
	return 0
}

func cw(a, b, c pt, include_collinear bool) bool {
	v := orientation(a, b, c)
	return v < 0 || include_collinear && v == 0
}

func collinear(a, b, c pt) bool {
	return orientation(a, b, c) == 0
}

func distance(a, b pt) int64 {
	dx := int64(a.x - b.x)
	dy := int64(a.y - b.y)
	return dx*dx + dy*dy
}

func convex_hull(pts []pt) []pt {

	p0 := pts[0]

	for i := 1; i < len(pts); i++ {
		if pts[i].x < p0.x || pts[i].x == p0.x && pts[i].y < p0.y {
			p0 = pts[i]
		}
	}

	sort.Slice(pts, func(i, j int) bool {
		v := orientation(p0, pts[i], pts[j])
		if v == 0 {
			return distance(p0, pts[i]) < distance(p0, pts[j])
		}
		return v < 0
	})

	n := len(pts)

	res := make([]pt, n)
	var p int

	for i := range n {
		for p > 1 && orientation(res[p-2], res[p-1], pts[i]) > 0 {
			p--
		}
		res[p] = pts[i]
		p++
	}

	return res[:p]
}
