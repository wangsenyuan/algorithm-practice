package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) float64 {
	pos := make(map[int]int)
	n := len(a)
	var sum int
	for i := range n {
		if j, ok := pos[a[i]]; ok {
			sum += (i-j)*(n-i)*2 - 1
		} else {
			sum += (i+1)*(n-i)*2 - 1
		}
		pos[a[i]] = i
	}

	tot := n * n

	return float64(sum) / float64(tot)
}
