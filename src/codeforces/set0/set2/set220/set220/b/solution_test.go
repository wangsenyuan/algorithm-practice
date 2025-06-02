package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))

	ans := process(reader)

	for _, x := range ans {
		expect := readNum(reader)
		if expect != x {
			t.Fatalf("Sample expect %d, but got %d", expect, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `7 2
3 1 2 2 3 3 7
1 7
3 4
3
1`)
}

func TestSample2(t *testing.T) {
	runSample(t, `6 6
1 2 2 3 3 3
1 2
2 2
1 3
2 4
4 6
1 6
1
0
2
1
1
3`)
}
