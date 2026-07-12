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
	fmt.Printf("%.10f %.10f %.10f\n", res[0], res[1], res[2])
}

func drive(reader *bufio.Reader) [3]float64 {
	var r, x1, y1, x2, y2 float64
	fmt.Fscan(reader, &r, &x1, &y1, &x2, &y2)
	return solve(r, x1, y1, x2, y2)
}

func solve(r, x1, y1, x2, y2 float64) [3]float64 {
	dx := x2 - x1
	dy := y2 - y1
	dist := math.Hypot(dx, dy)

	if dist >= r {
		return [3]float64{x1, y1, r}
	}

	if dist == 0 {
		return [3]float64{x1 + r/2, y1, r / 2}
	}

	nr := (r + dist) / 2
	ux := dx / dist
	uy := dy / dist
	return [3]float64{x2 - ux*nr, y2 - uy*nr, nr}
}
