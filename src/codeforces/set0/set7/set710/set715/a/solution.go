package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	res := solve(n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i := range n {
		fmt.Fprintln(writer, res[i])
	}
}

func solve(n int) []int {
	res := make([]int, n)
	// 1, 2, 3, 4, 5
	// 2, 2, 6,
	x := 2

	// 这个会溢出
	for i := range n {
		// 下一层时 i + 2
		w := (i + 2)
		j := 1
		if (w*w)%(i+1) != 0 {
			j = i + 1
		}
		res[i] = (w*w*j*j - x) / (i + 1)
		x = w * j
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
