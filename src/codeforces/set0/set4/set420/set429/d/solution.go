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
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(n, a)
}

type point struct {
	x int
	y int
}

func dist2(a point, b point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	return dx*dx + dy*dy
}

func solve(n int, a []int) int {
	arr1 := make([]point, n)
	var sum int
	for i := range n {
		sum += a[i]
		arr1[i] = point{i, sum}
	}

	// arr已经按照x升序了
	arr2 := slices.Clone(arr1)
	slices.SortFunc(arr2, func(a, b point) int {
		return cmp.Or(a.y-b.y, a.x-b.x)
	})
	// arr2 按照y升序

	var f func(x []point, y []point) int

	marked := make([]bool, n)

	bruteForce := func(x []point) int {
		d := 1 << 60

		for i := range x {
			for j := range i {
				d = min(d, dist2(x[i], x[j]))
			}
		}

		return d
	}

	merge := func(strip []point, d int) int {
		// strip is sorted by y asc
		for i := range strip {
			for j := i + 1; j < min(len(strip), i+8); j++ {
				d2 := dist2(strip[i], strip[j])
				d = min(d, d2)
			}
		}
		return d
	}

	f = func(x []point, y []point) int {
		if len(x) <= 3 {
			return bruteForce(x)
		}
		mid := len(x) / 2
		leftX := x[:mid]
		rightX := x[mid:]

		for i, p := range x {
			marked[p.x] = false
			if i < mid {
				marked[p.x] = true
			}
		}

		var leftY []point
		var rightY []point
		for _, p := range y {
			if marked[p.x] {
				leftY = append(leftY, p)
			} else {
				rightY = append(rightY, p)
			}
		}
		d1 := f(leftX, leftY)
		d2 := f(rightX, rightY)
		d := min(d1, d2)

		var strip []point
		for _, p := range y {
			dx := x[mid].x - p.x
			if dx*dx < d {
				strip = append(strip, p)
			}
		}

		d3 := merge(strip, d)

		return min(d1, d2, d3)
	}

	return f(arr1, arr2)
}

func abs(num int) int {
	return max(num, -num)
}
