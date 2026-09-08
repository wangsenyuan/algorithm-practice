package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		fmt.Fprintf(writer, "%.10f\n", drive(reader))
	}
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	c := make([]int, n)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &c[i], &p[i])
	}
	return solve(c, p)
}

func solve(c, p []int) float64 {
	n := len(c)
	var f float64
	for i := n - 1; i >= 0; i-- {
		f = max(f, f*(1.0-float64(p[i])/100)+float64(c[i]))
	}

	return f
}
