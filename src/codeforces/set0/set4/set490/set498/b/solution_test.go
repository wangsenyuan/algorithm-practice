package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)

	var expect float64
	fmt.Fscanf(reader, "%f", &expect)

	if math.Abs(res-expect) > 1e-6 {
		t.Errorf("Sample expect %f, but got %f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 2
50 2
10 1
1.5
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 2
0 2
100 2
1.0
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 3
50 3
50 2
25 2
1.687500000
`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `2 2
0 2
0 2
1
`
	runSample(t, s)
}
