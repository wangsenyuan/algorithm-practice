package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	for _, cur := range res {
		fmt.Println(cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) [][]int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	return solve(n, m)
}

func solve(n int, m int) [][]int {
	if n == 0 {
		return [][]int{{0, 1}, {0, m}, {0, 0}, {0, m - 1}}
	}
	if m == 0 {
		return [][]int{{1, 0}, {n, 0}, {0, 0}, {n - 1, 0}}
	}

	if n == 1 && m == 1 {
		return [][]int{{0, 0}, {1, 1}, {1, 0}, {0, 1}}
	}

	calc := func(arr [][]int) float64 {
		var dist float64
		for i := 0; i < 3; i++ {
			dx := arr[i][0] - arr[i+1][0]
			dy := arr[i][1] - arr[i+1][1]
			dist += math.Sqrt(float64(dx*dx + dy*dy))
		}

		return dist
	}

	var special = [][]int{
		{0, 0},
		{0, 1},
		{1, 0},
		{n, m},
		{n, m - 1},
		{n - 1, m},
	}
	if n > 1 {
		special = append(special, []int{n, 0})
	}
	if m > 1 {
		special = append(special, []int{0, m})
	}

	// len(special) = 8, C(8, 4) = 70 * 4! = 1680

	var res [][]int
	var sum float64

	var play func(mask int)

	var buf [][]int

	play = func(mask int) {
		if len(buf) == 4 {
			tmp := calc(buf)
			if tmp > sum {
				sum = tmp
				res = append([][]int{}, buf...)
			}
			return
		}
		for i := range len(special) {
			if (mask>>i)&1 == 0 {
				buf = append(buf, special[i])
				play(mask | (1 << i))
				buf = buf[:len(buf)-1]
			}
		}
	}
	play(0)

	return res
}
