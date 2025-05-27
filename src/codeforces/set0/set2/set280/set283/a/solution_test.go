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

	for i, x := range res {
		var y float64
		fmt.Fscanf(reader, "%f\n", &y)
		if math.Abs(x-y)/y > 1e-6 {
			t.Fatalf("Sample result %v, not correct at %d, expect %.7f	", res, i, y)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
2 1
3
2 3
2 1
3
0.500000
0.000000
1.500000
1.333333
1.500000
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `6
2 1
1 2 20
2 2
1 2 -3
3
3
0.500000
20.500000
14.333333
12.333333
17.500000
17.000000
`)
}
