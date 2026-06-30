package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.Abs(res-expect)/math.Max(1, math.Abs(expect)) > 1e-6 {
		t.Fatalf("Sample expect %.9f, but got %.9f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `10 0 1
-10 0 2 8
`, 9.584544103)
}

func TestSample2(t *testing.T) {
	runSample(t, `50 60 10
50 60 20 40
`, 0)
}

func TestTangentPointsOnOrbit(t *testing.T) {
	p1, p2 := tangentPoints(-10, 0, 10, 8)

	checkTangentPoint(t, p1, -10, 0, 10, 8)
	checkTangentPoint(t, p2, -10, 0, 10, 8)
	if math.Hypot(p1.x-p2.x, p1.y-p2.y) < 1e-8 {
		t.Fatalf("expect two different tangent points, got %+v and %+v", p1, p2)
	}
}

func checkTangentPoint(t *testing.T, p Point, x, y, R, r float64) {
	t.Helper()
	if math.Abs(math.Hypot(p.x, p.y)-R) > 1e-8 {
		t.Fatalf("point %+v is not on orbit radius %.9f", p, R)
	}
	dist := math.Abs(x*p.y-y*p.x) / math.Hypot(p.x-x, p.y-y)
	if math.Abs(dist-r) > 1e-8 {
		t.Fatalf("segment from (%.9f, %.9f) to %+v is not tangent to radius %.9f, distance %.9f", x, y, p, r, dist)
	}
}
