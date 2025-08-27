package main

import (
	"bytes"
	"fmt"
	"math"
)

func main() {
	var a, d float64
	fmt.Scan(&a, &d)
	var n int
	fmt.Scan(&n)
	res := solve(a, d, n)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintln(fmt.Sprintf("%.10f %.10f", x[0], x[1])))
	}

	fmt.Print(buf.String())
}

func solve(a float64, d float64, n int) [][]float64 {
	res := make([][]float64, n)

	var x, y float64
	pos := 0
	sum := 4 * a

	var move func(dist float64) (float64, float64)

	move = func(dist float64) (float64, float64) {
		if checkFloat(dist, 0) {
			return x, y
		}
		// try move to corner
		if pos == 0 {
			// 在第一条边
			if x+dist <= a {
				return x + dist, 0
			}
			dist -= a - x
			x = a
			pos++
		}
		if pos == 1 {
			// 在第二条边
			if y+dist <= a {
				return x, y + dist
			}
			dist -= a - y
			y = a
			pos++
		}
		if pos == 2 {
			if x-dist >= 0 {
				return x - dist, y
			}
			dist -= x
			x = 0
			pos++
		}

		if pos == 3 {
			if y-dist >= 0 {
				return x, y - dist
			}
			dist -= y
			y = 0
			pos = 0
		}

		return move(dist)
	}

	_, r := div(d, sum)

	for i := 0; i < n; i++ {
		// k is a integer
		nx, ny := move(r)

		res[i] = []float64{nx, ny}

		x, y = nx, ny
	}

	return res
}

func checkFloat(x float64, y float64) bool {
	return math.Abs(x-y) <= 1e-6
}

func div(a, b float64) (float64, float64) {
	x := float64(int(a / b))
	if b*x > a {
		x--
	}
	return x, a - x*b
}
