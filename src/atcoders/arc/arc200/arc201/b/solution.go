package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		buf.WriteString(fmt.Sprintf("%d\n", drive(reader)))
	}
	buf.WriteTo(os.Stdout)
}

func drive(reader *bufio.Reader) int {
	var n, W int
	fmt.Fscan(reader, &n, &W)
	X := make([]int, n)
	Y := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &X[i], &Y[i])
	}
	return solve(W, X, Y)
}

func solve(W int, X []int, Y []int) int {

	vals := make([][]int, 61)

	for i, v := range X {
		vals[v] = append(vals[v], Y[i])
	}

	var ans int

	for i := 0; i < 60; i++ {
		vals[i] = append(vals[i], 0, 0)
		sort.Ints(vals[i])
		if W&1 == 1 {
			ans += vals[i][len(vals[i])-1]
			vals[i] = vals[i][:len(vals[i])-1]
		}
		if len(vals[i])%2 == 1 {
			vals[i] = append(vals[i], 0)
		}
		sort.Ints(vals[i])
		for j := 0; j < len(vals[i]); j += 2 {
			vals[i+1] = append(vals[i+1], vals[i][j]+vals[i][j+1])
		}
		W >>= 1
	}

	return ans
}
