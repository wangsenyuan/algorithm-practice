package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	return solve(points)
}

type pair struct {
	first  int
	second int
}

func solve(points [][]int) int {
	n := len(points)
	res := n * (n - 1) * (n - 2) / 6

	var res2 int
	for i, a := range points {
		freq := make(map[pair]int)
		var vertical int
		var horizontal int
		for j := i + 1; j < n; j++ {
			dx := points[j][0] - a[0]
			dy := points[j][1] - a[1]
			if dx == 0 {
				// dy != 0
				res2 += vertical
				vertical++
			} else if dy == 0 {
				res2 += horizontal
				horizontal++
			} else {
				s := sign(dx) * sign(dy)
				c := gcd(abs(dx), abs(dy))
				k := pair{s * abs(dx) / c, abs(dy) / c}
				res2 += freq[k]
				freq[k]++
			}
		}
	}

	return res - res2
}

func sign(num int) int {
	if num == 0 {
		return 0
	}
	if num > 0 {
		return 1
	}
	return -1
}

func abs(num int) int {
	return max(num, -num)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
