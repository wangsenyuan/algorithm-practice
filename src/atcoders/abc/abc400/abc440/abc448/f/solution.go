package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) (points [][]int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	points = make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	res = solve(points)
	return
}

type data struct {
	id int
	x  int
	y  int
}

func solve(points [][]int) []int {
	n := len(points)
	arr := make([]data, n)
	for i, p := range points {
		arr[i] = data{id: i, x: p[0], y: p[1]}
	}

	blockSize := 2e7 / int(math.Sqrt(float64(n)))

	slices.SortFunc(arr, func(a, b data) int {
		if a.x/blockSize != b.x/blockSize {
			return a.x - b.x
		}
		p := a.x / blockSize
		if p&1 == 0 {
			return a.y - b.y
		}
		return b.y - a.y
	})

	res := make([]int, n)
	for i := range n {
		res[i] = arr[i].id + 1
	}
	for i := range n {
		if res[i] == 1 {
			slices.Reverse(res[:i+1])
			break
		}
	}
	return res
}
