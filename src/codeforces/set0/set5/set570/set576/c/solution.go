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
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	return solve(n, points)
}

type pt struct {
	id int
	x  int
	y  int
}

func solve(n int, points [][]int) []int {
	arr := make([]pt, n)
	for i := range n {
		arr[i] = pt{id: i, x: points[i][0], y: points[i][1]}
	}

	block_sz := 1000

	slices.SortFunc(arr, func(a, b pt) int {
		if a.x/block_sz != b.x/block_sz {
			return a.x/block_sz - b.x/block_sz
		}
		i := a.x / block_sz
		if i&1 == 0 {
			return a.y - b.y
		}
		return b.y - a.y
	})

	ans := make([]int, n)

	for i := range n {
		ans[i] = arr[i].id + 1
	}

	return ans
}
