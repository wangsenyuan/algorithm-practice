package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := drive(reader)
	fmt.Printf("%.10f\n", ans)
}

func drive(reader *bufio.Reader) float64 {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	queries := make([][]int, m)
	for i := range m {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(n, queries)
}

func solve(n int, queries [][]int) float64 {
	var pos int
	var neg int
	var sum int
	for _, cur := range queries {
		x, d := cur[0], cur[1]
		sum += x * n
		if d >= 0 {
			pos += d
		} else {
			neg += d
		}
	}
	for i := 1; i < n; i++ {
		sum += pos * i
	}
	// neg 始终是从中间开始算的
	mid := n / 2
	for i := range n {
		sum += neg * abs(i-mid)
	}

	return float64(sum) / float64(n)
}

func abs(num int) int {
	return max(num, -num)
}
