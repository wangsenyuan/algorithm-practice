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
	res := drive(reader)

	for _, x := range res {
		var y float64
		fmt.Fscan(reader, &y)
		if math.Abs(x-y)/max(1, y) > 1e-6 {
			t.Fatalf("Sample expect %.10f, but got %.10f", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 10 4 1
0 5
2 4
7 9
3
8
5
6.25
1.5
5`)
}

func TestSample2(t *testing.T) {
	runSample(t, `1 3 100 100 1
1 2
0
1
2
100
98.01
98`)
}
