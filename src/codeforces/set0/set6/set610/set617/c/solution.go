package main

import (
	"bufio"
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
	var n, x1, y1, x2, y2 int
	fmt.Fscan(reader, &n, &x1, &y1, &x2, &y2)
	flowers := make([][]int, n)
	for i := 0; i < n; i++ {
		flowers[i] = make([]int, 2)
		fmt.Fscan(reader, &flowers[i][0], &flowers[i][1])
	}
	return solve([]int{x1, y1}, []int{x2, y2}, flowers)
}

type point struct {
	x int
	y int
	d int
}

func solve(first []int, second []int, flowers [][]int) int {
	// 分组，离哪个近，就属于哪个组
	n := len(flowers)
	arr := make([]point, n)
	for i, cur := range flowers {
		arr[i] = point{cur[0], cur[1], distance(cur, first)}
	}

	slices.SortFunc(arr, func(a point, b point) int {
		return a.d - b.d
	})

	best := arr[n-1].d
	var suf int
	for i := n - 1; i >= 0; i-- {
		best = min(best, arr[i].d+suf)
		cur := distance([]int{arr[i].x, arr[i].y}, second)
		suf = max(suf, cur)
	}
	best = min(best, suf)

	return best
}

func distance(a []int, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}
