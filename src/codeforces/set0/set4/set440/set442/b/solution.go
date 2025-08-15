package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	p := make([]float64, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}

	return solve(p)
}

func solve(p []float64) float64 {
	x := slices.Max(p)
	if x == 1.0 {
		return 1.0
	}
	sort.Float64s(p)
	slices.Reverse(p)

	var res float64
	var S, P float64 = 0, 1
	for _, x := range p {
		S += x / (1 - x)
		P *= (1 - x)
		res = max(res, P*S)
	}

	return res
}
