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
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n, w, h int
	fmt.Fscan(reader, &n, &w, &h)
	dancers := make([][]int, n)
	for i := range n {
		var tp, pos, start int
		fmt.Fscan(reader, &tp, &pos, &start)
		dancers[i] = []int{tp, pos, start}
	}
	return solve(w, h, dancers)
}

type person struct {
	id  int
	pos int
	gp  int
}

func solve(w int, h int, dancers [][]int) [][]int {
	// 要算出所有碰撞点的时间和位置
	// 考虑x轴的p[i], 它从时刻t开始运动， 那么它需要那些y轴上的，t[j] + p[i]= t[i] + p[j]
	// p[j] - t[j] = p[i] - t[i]的会碰赚在一起

	// 假设有两个 (0, 10) 和 (10, 0), 它们同时出发，可以碰撞
	// 如果 (0, 5) 和 （10， 0) 那么第二个必须在时刻5开始出发， 才能碰撞

	n := len(dancers)

	var offset int

	for _, cur := range dancers {
		offset = max(offset, cur[1]-cur[2], cur[2]-cur[1])
	}

	s := make([][]int, 2*offset+1)

	for i, cur := range dancers {
		d := cur[1] - cur[2]
		s[d+offset] = append(s[d+offset], i)
	}
	ans1 := make([]int, n)
	ans2 := make([]int, n)

	for _, vs := range s {
		var xs []int
		var ys []int
		var arr []person
		for _, i := range vs {
			if dancers[i][0] == 1 {
				xs = append(xs, dancers[i][1])
			} else {
				ys = append(ys, dancers[i][1])
			}
			arr = append(arr, person{id: i, pos: dancers[i][1], gp: dancers[i][0]})
		}
		slices.Sort(xs)
		slices.Sort(ys)

		sort.Slice(arr, func(i int, j int) bool {
			if arr[i].gp != arr[j].gp {
				return arr[i].gp == 2
			}
			if arr[i].gp == 2 {
				return arr[i].pos > arr[j].pos
			}
			return arr[i].pos < arr[j].pos
		})

		for j := range len(xs) {
			ans1[arr[j].id] = xs[j]
			ans2[arr[j].id] = h
		}

		for j := range len(ys) {
			ans1[arr[j+len(xs)].id] = w
			ans2[arr[j+len(xs)].id] = ys[len(ys)-1-j]
		}
	}

	res := make([][]int, n)
	for i := range n {
		res[i] = []int{ans1[i], ans2[i]}
	}
	return res
}
