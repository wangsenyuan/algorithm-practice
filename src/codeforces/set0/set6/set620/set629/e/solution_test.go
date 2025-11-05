package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []float64) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)
	if len(ans) != len(expect) {
		t.Fatalf("Sample expect %v, but got %v", expect, ans)
	}

	for i, x := range expect {
		y := ans[i]
		if math.Abs(y-x) > 1e-6 {
			t.Fatalf("Sample expect %v, but got %v", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 3
2 4
4 1
3 2
3 1
2 3
4 1
`
	expect := []float64{4, 3, 3}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3 3
1 2
1 3
1 2
1 3
2 3
`
	expect := []float64{
		2.50000000,
		2.50000000,
		3.00000000}
	runSample(t, s, expect)
}
