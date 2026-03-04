package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	points, res := drive(reader)

	expect := make([][]int, len(res))
	for i := range len(res) {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		expect[i] = []int{a, b}
	}

	n := len(points)
	marked := make([]bool, n)
	calc := func(arr [][]int) int {
		clear(marked)

		var sum int

		for _, cur := range arr {
			a, b := cur[0]-1, cur[1]-1
			if marked[a] || marked[b] {
				continue
			}
			marked[a] = true
			marked[b] = true
			sum += abs(points[a][0]-points[b][0]) + abs(points[a][1]-points[b][1])
		}
		return sum
	}

	a := calc(res)
	b := calc(expect)

	if a != b {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestSample1(t *testing.T) {
	s := `4
1 1
3 0
4 2
3 4
4 1
2 3
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10
-1 -1
-1 2
-2 -2
-2 0
0 2
2 -3
-4 -4
-4 -2
0 1
-4 -2
8 1
9 10
7 5
2 3
6 4
`
	runSample(t, s)
}
