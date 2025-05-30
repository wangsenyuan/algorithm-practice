package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for _, x := range res {
		y := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5 5
1 3 1 2 4
2 1 4
2 1 5
2 2 4
1 3 10
2 1 5
12
32
8
50
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `5 4
1 3 1 2 4
3 1 4 1
2 2 4
1 2 10
2 1 5
12
45
`)
}
