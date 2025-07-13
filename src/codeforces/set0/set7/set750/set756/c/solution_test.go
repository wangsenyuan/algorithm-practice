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
			t.Fatalf("Sample expect %s, but got %v", s, res)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `2
2 1 2
1 0
2
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 1 2
2 1 3
3 0
2
3
2`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
5 0
4 0
3 1 1
2 1 1
1 1 2
-1
-1
-1
-1
2`)
}
