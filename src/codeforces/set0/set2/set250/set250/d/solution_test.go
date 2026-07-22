package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, input string, expect []int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(input))
	res, a, b, ya, yb, l := drive(reader)

	play := func(i int, j int) float64 {
		i--
		j--

		d1 := math.Sqrt(float64(a)*float64(a) + float64(ya[i])*float64(ya[i]))
		d2 := math.Sqrt(float64(b-a)*float64(b-a) + float64(yb[j]-ya[i])*float64(yb[j]-ya[i]))

		return d1 + d2 + float64(l[j])
	}

	expectAns := play(expect[0], expect[1])
	gotAns := play(res[0], res[1])
	if math.Abs(expectAns-gotAns) > 1e-6 {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 2 3 5
-2 -1 4
-1 2
7 3
`, []int{2, 2})
}
