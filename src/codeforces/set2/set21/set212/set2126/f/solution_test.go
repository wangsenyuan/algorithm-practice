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
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `1 1
1
1 1
0`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `2 3
1 1
1 2 10
1 2
2 2
1 1
10
0
10`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `5 4
1 2 1 2 3
1 2 5
2 3 3
2 4 4
4 5 7
3 2
5 2
1 2
2 3
12
5
0
12`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `4 3
1 1 2 2
1 2 2
2 3 6
2 4 8
3 1
4 1
2 2
8
0
16`
	runSample(t, s)
}
