package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			fmt.Fscan(reader, &points[i][j])
		}
	}
	return solve(points)
}

func solve(points [][]int) []int {

	getVector := func(i int, j int) []int {
		return []int{points[j][0] - points[i][0], points[j][1] - points[i][1], points[j][2] - points[i][2], points[j][3] - points[i][3], points[j][4] - points[i][4]}
	}
	getScalarProduct := func(x []int, y []int) int {
		return x[0]*y[0] + x[1]*y[1] + x[2]*y[2] + x[3]*y[3] + x[4]*y[4]
	}

	n := len(points)

	check := func(i int) bool {
		if n > 12 {
			return true
		}
		for j := 0; j < n ; j++ {
			if i == j {
				continue
			}
			for k := range j {
				if k == i {
					continue
				}
				x := getVector(i, j)
				y := getVector(i, k)
				s := getScalarProduct(x, y)
				if s > 0 {
					return true
				}
			}
		}
		return false
	}

	var res []int
	for i := range n {
		bad := check(i)
		if !bad {
			res = append(res, i+1)
		}
	}
	return res
}
