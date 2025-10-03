package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	points, res := drive(reader)

	check := func(i int, j int) float64 {
		i--
		j--
		return nonOrientedAngle(points[i][0], points[i][1], points[j][0], points[j][1])
	}

	a := check(res[0], res[1])
	b := check(res[1], res[0])
	if math.Abs(a-b) > 1e-6 {
		t.Fatalf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
-1 0
0 -1
1 0
1 1
`
	expect := []int{3, 4}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `6
-1 0
0 -1
1 0
1 1
-4 -5
-4 -6
`
	expect := []int{6, 5}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
-5120 -3251
8269 -7984
841 3396
3136 -7551
-1280 -3013
-3263 -3278
`
	expect := []int{6, 1}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4
9800 9981
61 9899
-9926 -9932
-149 -9926
`
	expect := []int{3, 4}
	runSample(t, s, expect)
}
