package main

import (
	"bufio"
	"strings"
	"testing"
)

type segment struct {
	x0, y0, x1, y1 int
}

func (s segment) intersect(t segment) bool {
	return s.x0 <= t.x1 && t.x0 <= s.x1 && s.y0 <= t.y1 && t.y0 <= s.y1
}

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	n, edges, points, res := process(reader)

	pos := make([]int, n)
	for i, j := range res {
		pos[j-1] = i
	}

	getSegment := func(u int, v int) segment {
		i := pos[u]
		j := pos[v]
		return segment{
			x0: points[i][0],
			y0: points[i][1],
			x1: points[j][0],
			y1: points[j][1],
		}
	}

	checkIntersect := func(x, y []int) bool {
		s1 := getSegment(x[0]-1, x[1]-1)
		s2 := getSegment(y[0]-1, y[1]-1)
		return s1.intersect(s2)
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n-1; j++ {
			x := edges[i]
			y := edges[j]

			if x[0] == y[0] || x[0] == y[1] || x[1] == y[0] || x[1] == y[1] {
				continue
			}

			if checkIntersect(x, y) {
				t.Errorf("edges %d and %d cross", i+1, j+1)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	sample := `3
1 3
2 3
0 0
1 1
2 0
`
	runSample(t, sample)
}

func TestSample2(t *testing.T) {
	sample := `4
1 2
2 3
1 4
-1 -2
3 5
-3 3
2 0
`
	runSample(t, sample)
}

func TestSample3(t *testing.T) {
	sample := `9
9 8
6 1
4 2
1 8
8 7
5 7
6 2
3 2
8 3
-5 -4
-6 4
6 5
5 -1
1 -8
9 -3
5 5
8 1
`
	runSample(t, sample)
}
