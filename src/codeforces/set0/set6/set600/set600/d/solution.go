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
	fmt.Printf("%.10f\n", res)
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) float64 {
	first := readNNums(reader, 3)
	second := readNNums(reader, 3)
	return solve(first, second)
}

func solve(first []int, second []int) float64 {
	if first[2] > second[2] {
		first, second = second, first
	}
	dx := first[0] - second[0]
	dy := first[1] - second[1]
	r1 := float64(first[2])
	r2 := float64(second[2])
	dist := math.Sqrt(float64(dx*dx + dy*dy))

	if dist >= r1+r2 {
		// first在second的外部
		return 0
	}

	if dist <= r2-r1 {
		return math.Pi * r1 * r1
	}

	return solve2(r1, r2, dist) + solve2(r2, r1, dist)
}

func solve2(r1 float64, r2 float64, d float64) float64 {
	alpha := math.Acos((r2*r2 + d*d - r1*r1) / (2 * r2 * d))
	s := r2 * r2 * alpha

	t := r2 * r2 * math.Sin(alpha) * math.Cos(alpha)

	return s - t
}
