package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1:len(s)-1])
}

func drive(reader *bufio.Reader) (points [][]int, l int, res []int) {
	var n int
	fmt.Fscan(reader, &n, &l)
	points = make([][]int, n+1)
	points[0] = []int{0, 0}
	for i := 1; i <= n; i++ {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}

	res = solve(l, points)
	points = points[1:]
	return
}

func solve(l int, points [][]int) []int {
	n := len(points)
	fa := make([]int, n)

	dp := make([]float64, n)

	cost := func(i int, j int) float64 {
		dist := float64(points[j][0] - points[i][0])
		return math.Sqrt(math.Abs(dist - float64(l)))
	}

	check := func(r float64) bool {
		for i := range n {
			dp[i] = math.MaxFloat64
			fa[i] = -1
		}
		dp[0] = 0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				// cost(i, j) - r * b[j]
				tmp := dp[i] + cost(i, j) - r*float64(points[j][1])
				if tmp < dp[j] {
					dp[j] = tmp
					fa[j] = i
				}
			}
		}
		return dp[n-1] < 0
	}

	var lo, hi float64 = 0, 1e12

	for range 100 {
		mid := (lo + hi) / 2
		if check(mid) {
			hi = mid
		} else {
			lo = mid
		}
	}

	check((lo + hi) / 2)

	var res []int
	for i := n - 1; i > 0; i = fa[i] {
		res = append(res, i)
	}

	slices.Reverse(res)
	return res
}
