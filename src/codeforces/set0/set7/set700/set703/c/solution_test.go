package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if math.Abs(res-expect)/max(1, expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5 1 2
1 2
3 1
4 3
3 4
1 4
`, 5.0)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 10 2 4
-4 5
-3 2
-1 0
3 0
5 2
6 5
5 8
3 10
-1 10
-2 9
`, 4.5)
}
