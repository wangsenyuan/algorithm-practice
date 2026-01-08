package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for _, x := range res {
		var y int
		fmt.Fscan(reader, &y)
		if x != y {
			t.Errorf("Sample expect %v, but got %v", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 3
8 -5 -4 1 -7 4
0 7 3
2
0
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10 5
-1 -1 -1 -1 -1 -1 -1 -1 -1 -1
0 1 2 3 4
10
9
8
7
6`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1 5
-7
1 3 5 7 9
1
1
1
0
0`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `5 5
0 0 -1 0 1
1 0 1 5 8
0
1
0
0
0`
	runSample(t, s)
}
