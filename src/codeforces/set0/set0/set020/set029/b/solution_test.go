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

	if math.Abs(expect-res) > 1e-7 {
		t.Fatalf("Sample expect %.7f but got %.7f", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 1 3 4 5
0.66666667
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5 4 3 1 1
2.33333333
`
	runSample(t, s)
}

