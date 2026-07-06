package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%.8f\n", drive(reader))
}

func drive(reader *bufio.Reader) float64 {
	pts := make([][]float64, 3)
	for i := range 3 {
		pts[i] = make([]float64, 2)
		fmt.Fscan(reader, &pts[i][0], &pts[i][1])
	}
	return solve(pts)
}

func solve(pts [][]float64) float64 {
	o := circumcenter(pts[0], pts[1], pts[2])
	r := dist(o, pts[0])

	angles := make([]float64, 3)
	for i, p := range pts {
		angles[i] = math.Atan2(p[1]-o[1], p[0]-o[0])
		if angles[i] < 0 {
			angles[i] += 2 * math.Pi
		}
	}
	sort.Float64s(angles)

	arcs := []float64{
		angles[1] - angles[0],
		angles[2] - angles[1],
		2*math.Pi - angles[2] + angles[0],
	}

	for n := 3; n <= 100; n++ {
		step := 2 * math.Pi / float64(n)
		ok := true
		for _, arc := range arcs {
			x := arc / step
			if math.Abs(x-math.Round(x)) > 1e-5 {
				ok = false
				break
			}
		}
		if ok {
			return float64(n) * r * r * math.Sin(step) / 2
		}
	}

	return 0
}

func circumcenter(a, b, c []float64) []float64 {
	d := 2 * (a[0]*(b[1]-c[1]) + b[0]*(c[1]-a[1]) + c[0]*(a[1]-b[1]))
	aa := a[0]*a[0] + a[1]*a[1]
	bb := b[0]*b[0] + b[1]*b[1]
	cc := c[0]*c[0] + c[1]*c[1]
	x := (aa*(b[1]-c[1]) + bb*(c[1]-a[1]) + cc*(a[1]-b[1])) / d
	y := (aa*(c[0]-b[0]) + bb*(a[0]-c[0]) + cc*(b[0]-a[0])) / d
	return []float64{x, y}
}

func dist(a, b []float64) float64 {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return math.Sqrt(dx*dx + dy*dy)
}
