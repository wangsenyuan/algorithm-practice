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
	fmt.Printf("%.10f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, l, r int
	fmt.Fscan(reader, &n, &l, &r)
	lights := make([][]int, n)
	for i := range n {
		lights[i] = make([]int, 3)
		for j := range 3 {
			fmt.Fscan(reader, &lights[i][j])
		}
	}
	return solve(l, r, lights)
}

const inf = 1 << 60

func solve(l int, r int, lights [][]int) float64 {

	play := func(x int, y int, alpha float64, x0 float64) (x1 float64) {
		// 一个在位置 v (x, y)处,可以照射最大角度为a, 在ox轴上的投影,左端点为x0时
		// 右端点y0 = ?
		// 假设 y0 - x0 = w
		// w * y / 2 = lengh(v, x0) * length(v, y0) * sin(a)
		d1 := math.Abs(float64(x) - x0)
		// tan(beata) = d1 / y
		beata := math.Atan(d1 / float64(y))

		halfPi := math.Pi / 2

		if float64(x) < x0 {
			beata += alpha
			if beata >= halfPi {
				// ray goes above horizontal, illuminates all x-axis to the right
				return inf
			} else {
				w := float64(y) * math.Tan(beata)
				x1 = float64(x) + w
			}
		} else {
			if beata < alpha {
				beata = alpha - beata
				w := float64(y) * math.Tan(beata)
				x1 = float64(x) + w
			} else {
				beata -= alpha
				w := float64(y) * math.Tan(beata)
				x1 = float64(x) - w
			}
		}
		return
	}
	n := len(lights)
	dp := make([]float64, 1<<n)
	for i := range 1 << n {
		dp[i] = -inf
	}
	dp[0] = float64(l)

	for mask := range 1 << n {
		if dp[mask] > -inf {
			x0 := dp[mask]
			for i := range n {
				if (mask>>i)&1 == 0 {
					x1 := play(lights[i][0], lights[i][1], getAngle(lights[i][2]), x0)
					if x1 >= float64(r) {
						return float64(r - l)
					}
					dp[mask|1<<i] = max(dp[mask|1<<i], x1)
				}
			}
		}
	}
	ans := min(dp[(1<<n)-1], float64(r))
	return ans - float64(l)
}

func getAngle(a int) float64 {
	// pi = 180
	return math.Pi * float64(a) / 180
}
