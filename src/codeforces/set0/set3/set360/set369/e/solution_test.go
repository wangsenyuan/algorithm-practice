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
	s := `3 3
1 3
4 5
6 7
3 1 4 7
2 4 5
1 8
3
1
0
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1 1
172921 894619
1 14141
0
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3 1
439010 864662
377278 743032
771051 955458
1 568232
2
`
	runSample(t, s)
}
