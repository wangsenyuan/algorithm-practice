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
	fmt.Printf("%.9f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, x, y int
	fmt.Fscan(reader, &n, &x, &y)
	polygon := make([][]int, n)
	for i := 0; i < n; i++ {
		polygon[i] = make([]int, 2)
		fmt.Fscan(reader, &polygon[i][0], &polygon[i][1])
	}
	return solve([]int{x, y}, polygon)
}

func solve(P []int, polygon [][]int) float64 {
	n := len(polygon)
	var r2 int
	arr := make([]Point, n)
	for i, cur := range polygon {
		r2 = max(r2, dist(P, cur))
		arr[i] = Point{cur[0], cur[1], i}
	}

	var r1 = math.Sqrt(float64(r2))

	O := Point{P[0], P[1], -1}

	for i := 0; i < len(arr); i++ {
		j := (i + 1) % len(arr)
		d := O.distance(arr[i], arr[j])
		r1 = min(r1, d)
	}

	r1 *= r1

	return (float64(r2) - r1) * math.Pi
}

func dist(a, b []int) int {
	dx := a[0] - b[0]
	dy := a[1] - b[1]
	return dx*dx + dy*dy
}

type Point struct {
	x, y int
	id   int
}

func (this Point) distance(a, b Point) float64 {
	// Vector from a to b
	abx := float64(b.x - a.x)
	aby := float64(b.y - a.y)

	// Vector from a to this
	apx := float64(this.x - a.x)
	apy := float64(this.y - a.y)

	// Calculate the dot product and squared length of ab
	ab_dot_ap := abx*apx + aby*apy
	ab_dot_ab := abx*abx + aby*aby

	// If a and b are the same point, return distance to a
	if ab_dot_ab == 0 {
		return math.Sqrt(apx*apx + apy*apy)
	}

	// Calculate the parameter t (projection of ap onto ab)
	t := ab_dot_ap / ab_dot_ab

	// If projection falls before a, return distance to a
	if t < 0 {
		return math.Sqrt(apx*apx + apy*apy)
	}

	// If projection falls after b, return distance to b
	if t > 1 {
		bpx := float64(this.x - b.x)
		bpy := float64(this.y - b.y)
		return math.Sqrt(bpx*bpx + bpy*bpy)
	}

	// Projection falls on the segment, calculate perpendicular distance
	projx := float64(a.x) + t*abx
	projy := float64(a.y) + t*aby

	dx := float64(this.x) - projx
	dy := float64(this.y) - projy
	return math.Sqrt(dx*dx + dy*dy)
}
