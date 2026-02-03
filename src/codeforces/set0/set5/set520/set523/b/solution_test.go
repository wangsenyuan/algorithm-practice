package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect [][]float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	for i, cur := range res {
		for j, x := range cur {
			y := expect[i][j]
			if math.Abs(x-y)/max(1, y) > 1e-4 {
				t.Fatalf("Sample expect %v, but got %v", expect, res)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `1 1 2.000000
1
1
1
`
	expect := [][]float64{
		{1.0, 0.5, 0.5},
	}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `11 4 1.250000
9 11 7 5 15 6 6 6 6 6 6
8
4 5 6 7 8 9 10 11
`
	expect := [][]float64{
		{8.000000, 4.449600, 0.443800},
		{9.500000, 6.559680, 0.309507},
		{8.250000, 6.447744, 0.218455},
		{8.000000, 6.358195, 0.205226},
		{8.250000, 6.286556, 0.237993},
		{6.000000, 6.229245, 0.038207},
		{6.000000, 6.183396, 0.030566},
		{6.000000, 6.146717, 0.024453},
	}
	runSample(t, s, expect)
}


func TestSample3(t *testing.T) {
	s := `13 4 1.250000
3 3 3 3 3 20 3 3 3 3 3 3 3
10
4 5 6 7 8 9 10 11 12 13
`
	expect := [][]float64{
		{3.000000, 1.771200, 0.409600},
		{3.000000, 2.016960, 0.327680},
		{7.250000, 5.613568, 0.225715},
		{7.250000, 5.090854, 0.297813},
		{7.250000, 4.672684, 0.355492},
		{7.250000, 4.338147, 0.401635},
		{3.000000, 4.070517, 0.356839},
		{3.000000, 3.856414, 0.285471},
		{3.000000, 3.685131, 0.228377},
		{3.000000, 3.548105, 0.182702},
	}
	runSample(t, s, expect)
}
