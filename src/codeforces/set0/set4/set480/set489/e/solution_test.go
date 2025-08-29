package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	points, l, res := drive(reader)

	get := func(arr []int) float64 {
		var frustration float64
		var sum int
		var prev int
		for i := 0; i < len(arr); i++ {
			dist := float64(points[arr[i]-1][0] - prev)
			frustration += math.Sqrt(math.Abs(dist - float64(l)))
			sum += points[arr[i]-1][1]
			prev = points[arr[i]-1][0]
		}
		return frustration / float64(sum)
	}

	x := get(expect)
	y := get(res)

	if y-x > 1e-7 {
		t.Errorf("expect %v(%.10f), but got %v(%.10f)", expect, x, res, y)
	}
}

func TestSample1(t *testing.T) {
	s := `5 9
10 10
20 10
30 1
31 5
40 10
	`
	runSample(t, s, []int{1, 2, 4, 5})
}
