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
	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 4 3
1 2 2
2 4 1
1 3 1
3 4 2
`
	expect := 1.5
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 11 23
1 2 3
2 3 4
3 4 5
4 5 6
1 3 4
2 4 5
3 5 6
1 4 2
2 5 3
1 5 2
3 2 30
`
	expect := 10.2222222222
	runSample(t, s, expect)
}
