package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1], res[2])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	stars := make([][]int, n)
	for i := range n {
		stars[i] = make([]int, 2)
		fmt.Fscan(reader, &stars[i][0], &stars[i][1])
	}
	return solve(n, stars)
}

type point struct {
	id int
	x  int
	y  int
}

func cross(a, b, c point) int {
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
}

const inf = 1 << 60

func dist2(a, b point) int {
	dx := a.x - b.x
	dy := a.y - b.y
	return dx*dx + dy*dy
}

func solve(n int, stars [][]int) []int {
	arr := make([]point, n)
	first := -1
	for i, cur := range stars {
		arr[i] = point{id: i, x: cur[0], y: cur[1]}
		if first == -1 || arr[first].x > arr[i].x {
			first = i
		}
	}

	second := -1
	for i := range n {
		if i == first {
			continue
		}
		if second == -1 || dist2(arr[first], arr[i]) < dist2(arr[first], arr[second]) {
			second = i
		}
	}

	third := -1
	best := inf
	for i := range n {
		if i == first || i == second {
			continue
		}
		tmp := cross(arr[first], arr[second], arr[i])
		if tmp == 0 {
			continue
		}
		if third == -1 || abs(tmp) < best {
			best = abs(tmp)
			third = i
		}
	}

	return []int{first + 1, second + 1, third + 1}
}

func abs(num int) int {
	return max(num, -num)
}
