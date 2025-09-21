package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a)
}

func solve(a []int) float64 {
	sort.Ints(a)
	n := len(a)
	mx := a[n-1] - a[0]
	dp := make([]float64, mx+1)

	// 每一对被选中的概率 = 1 / w
	w := n * (n - 1) / 2

	for i := range n {
		for j := range i {
			dp[a[i]-a[j]] += 1 / float64(w)
		}
	}

	dp2 := make([]float64, mx+1)
	for x := 1; x <= mx; x++ {
		for y := 1; x+y <= mx; y++ {
			dp2[x+y] += dp[x] * dp[y]
		}
	}
	var res float64

	for x := 1; x <= mx; x++ {
		for y := x + 1; y <= mx; y++ {
			res += dp2[x] * dp[y]
		}
	}
	return res
}
