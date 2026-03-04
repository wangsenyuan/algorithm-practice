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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, res := drive(reader)
		for _, cur := range res {
			fmt.Fprintln(writer, cur[0], cur[1])
		}
	}
}

func drive(reader *bufio.Reader) (points [][]int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n)
	points = make([][]int, n)
	for i := range n {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		points[i] = []int{x, y}
	}
	res = solve(points)
	return
}

type point struct {
	id int
	x  int
	y  int
}

func solve(points [][]int) [][]int {
	n := len(points)
	arr := make([]point, n)
	for i := 0; i < n; i++ {
		arr[i] = point{i, points[i][0], points[i][1]}
	}
	slices.SortFunc(arr, func(a, b point) int {
		return cmp.Or(a.x-b.x, a.y-b.y)
	})
	h := n / 2
	first := arr[:h]
	second := arr[h:]
	slices.SortFunc(first, func(a, b point) int {
		return a.y - b.y
	})
	slices.SortFunc(second, func(a, b point) int {
		return b.y - a.y
	})
	res := make([][]int, h)
	for i := range h {
		res[i] = []int{first[i].id + 1, second[i].id + 1}
	}
	return res
}
