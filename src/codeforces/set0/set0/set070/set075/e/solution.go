package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintf(writer, "%.10f\n", drive(reader))
}

func drive(reader *bufio.Reader) float64 {
	var xStart, yStart, xEnd, yEnd, n int
	fmt.Fscan(reader, &xStart, &yStart, &xEnd, &yEnd, &n)
	poly := make([][]int, n)
	for i := range n {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		poly[i] = []int{x, y}
	}
	return solve(xStart, yStart, xEnd, yEnd, poly)
}

type point struct {
	x float64
	y float64
}

func solve(xStart, yStart, xEnd, yEnd int, poly [][]int) float64 {
	s := point{float64(xStart), float64(yStart)}
	e := point{float64(xEnd), float64(yEnd)}
	p := make([]point, len(poly))
	for i, v := range poly {
		p[i] = point{float64(v[0]), float64(v[1])}
	}

	lo, hi, miss := clip(p, s, e)
	if miss || hi-lo <= eps {
		return dist(s, e)
	}
	a := at(s, e, lo)
	b := at(s, e, hi)
	if sameEdge(p, a, b) {
		return dist(s, e)
	}

	inside := dist(a, b)
	along := chain(p, a, b)
	return dist(s, e) - inside + min(2*inside, along, perimeter(p)-along)
}

const eps = 1e-9

func clip(p []point, s, e point) (lo, hi float64, miss bool) {
	lo, hi = 0, 1
	for i, a := range p {
		b := p[(i+1)%len(p)]
		cs := cross(a, b, s)
		ce := cross(a, b, e)
		if cs >= 0 && ce >= 0 {
			continue
		}
		if cs < 0 && ce < 0 {
			return 0, 0, true
		}
		t := cs / (cs - ce)
		if cs < 0 {
			lo = max(lo, t)
		} else {
			hi = min(hi, t)
		}
	}
	return lo, hi, lo > hi
}

func chain(p []point, from, to point) float64 {
	n := len(p)
	e := locate(p, from)
	goal := locate(p, to)
	cur := from
	var length float64
	for range n {
		if same(cur, to) {
			return length
		}
		end := p[(e+1)%n]
		if e == goal && onSeg(to, cur, end) {
			return length + dist(cur, to)
		}
		length += dist(cur, end)
		cur = end
		e = (e + 1) % n
	}
	return length
}

func locate(p []point, q point) int {
	n := len(p)
	for i := range n {
		v := p[(i+1)%n]
		if same(q, v) {
			continue
		}
		if onSeg(q, p[i], v) {
			return i
		}
	}
	return 0
}

func sameEdge(p []point, a, b point) bool {
	for i, u := range p {
		v := p[(i+1)%len(p)]
		if onSeg(a, u, v) && onSeg(b, u, v) {
			return true
		}
	}
	return false
}

func perimeter(p []point) float64 {
	var sum float64
	for i, u := range p {
		sum += dist(u, p[(i+1)%len(p)])
	}
	return sum
}

func at(s, e point, t float64) point {
	return point{s.x + (e.x-s.x)*t, s.y + (e.y-s.y)*t}
}

func cross(a, b, c point) float64 {
	return (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
}

func dist(a, b point) float64 {
	return math.Hypot(a.x-b.x, a.y-b.y)
}

func same(a, b point) bool {
	return dist(a, b) <= eps
}

func onSeg(q, a, b point) bool {
	if math.Abs(cross(a, b, q)) > eps {
		return false
	}
	return q.x >= min(a.x, b.x)-eps && q.x <= max(a.x, b.x)+eps &&
		q.y >= min(a.y, b.y)-eps && q.y <= max(a.y, b.y)+eps
}
