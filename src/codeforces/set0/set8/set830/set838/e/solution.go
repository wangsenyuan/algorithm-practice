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
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n int
	fmt.Fscan(reader, &n)
	points := make([][]int, n)
	for i := range n {
		points[i] = make([]int, 2)
		fmt.Fscan(reader, &points[i][0], &points[i][1])
	}
	return solve(points)
}

const inf = 1e18

func distance(p1 []int, p2 []int) float64 {
	dx := float64(p1[0] - p2[0])
	dy := float64(p1[1] - p2[1])
	return math.Sqrt(dx*dx + dy*dy)
}

func solve(points [][]int) float64 {
	n := len(points)

	dp := make([][][2]float64, n)
	for i := range n {
		dp[i] = make([][2]float64, n)
		tmp := distance(points[i%n], points[(i+1)%n])
		dp[i][(i+1)%n][0] = tmp
		dp[i][(i+1)%n][1] = tmp
	}

	// dp[l][r][0] 表示终点在l处, dp[l][r][1]表示终点在r处
	for d := 1; d < n; d++ {
		for i := range n {
			j := (i + d) % n
			i1 := (i + 1) % n
			j1 := (j - 1 + n) % n
			dp[i][j][0] = max(dp[i1][j][0]+distance(points[i], points[i1]), dp[i1][j][1]+distance(points[i], points[j]))
			dp[i][j][1] = max(dp[i][j1][0]+distance(points[i], points[j]), dp[i][j1][1]+distance(points[j], points[j1]))
		}
	}
	var res float64
	for i := range n {
		res = max(res, dp[i][(i-1+n)%n][0], dp[i][(i-1+n)%n][1])
	}

	return res
}
