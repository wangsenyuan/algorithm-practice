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
	s := `3 0 0
0 1
-1 2
1 2
`
	expect := 12.566370614359172464
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 1 -1
0 0
1 2
2 0
1 1
`
	expect := 21.991148575128551812
	runSample(t, s, expect)
}
