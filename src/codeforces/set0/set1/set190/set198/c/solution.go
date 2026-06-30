package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%.9f\n", drive(reader))
}

func drive(reader *bufio.Reader) float64 {
	var xp, yp, vp int
	fmt.Fscan(reader, &xp, &yp, &vp)
	var x, y, v, r int
	fmt.Fscan(reader, &x, &y, &v, &r)
	return solve(xp, yp, vp, x, y, v, r)
}

func solve(xp, yp, vp, x, y, v, r int) float64 {
	R := math.Hypot(float64(xp), float64(yp))
	start := Point{float64(x), float64(y)}
	speed := float64(v)
	rr := float64(r)
	a0 := math.Atan2(float64(yp), float64(xp))
	w := float64(vp) / R

	check := func(t float64) bool {
		a := a0 + w*t
		target := Point{R * math.Cos(a), R * math.Sin(a)}
		return shortestPath(start, target, rr)/speed <= t
	}

	var l, h float64 = 0, 1
	for !check(h) {
		h *= 2
	}

	for range 100 {
		mid := (l + h) / 2
		if check(mid) {
			h = mid
		} else {
			l = mid
		}
	}

	return h
}

type Point struct {
	x float64
	y float64
}

func shortestPath(a, b Point, r float64) float64 {
	if distanceToSegment(a, b) >= r-1e-10 {
		return dist(a, b)
	}

	ta := innerTangents(a, r)
	tb := innerTangents(b, r)
	best := math.Pi * 2
	for _, x := range ta {
		for _, y := range tb {
			best = math.Min(best, angleDiff(angle(x), angle(y)))
		}
	}

	return math.Sqrt(norm2(a)-r*r) + math.Sqrt(norm2(b)-r*r) + r*best
}

func innerTangents(p Point, r float64) []Point {
	d2 := norm2(p)
	h := math.Sqrt(d2 - r*r)

	a := r * r / d2
	b := r * h / d2

	return []Point{
		{
			a*p.x - b*p.y,
			a*p.y + b*p.x,
		},
		{
			a*p.x + b*p.y,
			a*p.y - b*p.x,
		},
	}
}

// (x, y)在圆(0, 0, r)的外部, 返回大圆R上的两个点, 从(x, y)连过去刚好与小圆r相切
func tangentPoints(x, y, R, r float64) (Point, Point) {
	ps := innerTangents(Point{x, y}, r)
	p1, p2 := ps[0], ps[1]

	return intersectOuter(Point{x, y}, p1, R), intersectOuter(Point{x, y}, p2, R)
}

func intersectOuter(start Point, touch Point, R float64) Point {
	dx := touch.x - start.x
	dy := touch.y - start.y

	A := dx*dx + dy*dy
	B := 2 * (start.x*dx + start.y*dy)
	C := start.x*start.x + start.y*start.y - R*R
	D := math.Sqrt(math.Max(0, B*B-4*A*C))

	t1 := (-B - D) / (2 * A)
	t2 := (-B + D) / (2 * A)
	t := math.Max(t1, t2)

	return Point{
		start.x + t*dx,
		start.y + t*dy,
	}
}

func distanceToSegment(a, b Point) float64 {
	dx := b.x - a.x
	dy := b.y - a.y
	d2 := dx*dx + dy*dy
	t := -(a.x*dx + a.y*dy) / d2
	t = math.Max(0, math.Min(1, t))
	x := a.x + t*dx
	y := a.y + t*dy
	return math.Hypot(x, y)
}

func dist(a, b Point) float64 {
	return math.Hypot(a.x-b.x, a.y-b.y)
}

func norm2(a Point) float64 {
	return a.x*a.x + a.y*a.y
}

func angle(a Point) float64 {
	return math.Atan2(a.y, a.x)
}

func angleDiff(a, b float64) float64 {
	d := math.Abs(a - b)
	if d > math.Pi {
		d = 2*math.Pi - d
	}
	return d
}
