package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.10f\n", res)
}

func process(reader *bufio.Reader) float64 {
	var n, x int
	fmt.Fscan(reader, &n, &x)
	a := make([]float64, x+1)
	for i := range x + 1 {
		fmt.Fscan(reader, &a[i])
	}

	return solve(n, a)
}

func solve(n int, a []float64) float64 {
	n--

	var id matrix

	for i := range 128 {
		for j, v := range a {
			u := i ^ j
			id[i][u] += v
		}
	}

	var pw [31]matrix
	pw[0] = id
	for i := 1; i < 31; i++ {
		pw[i] = pw[i-1].mul(pw[i-1])
	}

	var ans matrix
	for i := range 128 {
		ans[i][i] = 1
	}

	for i := 30; i >= 0; i-- {
		if (n>>i)&1 == 1 {
			ans = ans.mul(pw[i])
		}
	}

	var res matrix
	for i, v := range a {
		res[0][i] = v
	}

	fans := res.mul(ans)

	return 1.0 - fans[0][0]
}

type matrix [128][128]float64

func (a matrix) mul(b matrix) matrix {
	var c matrix
	for i := range 128 {
		for j := range 128 {
			for k := range 128 {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
