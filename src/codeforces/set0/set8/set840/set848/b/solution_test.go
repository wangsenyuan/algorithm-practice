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
	for _, cur := range res {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		if len(cur) == 0 || cur[0] != x || cur[1] != y {
			t.Fatalf("Sample expect %v, but got %v", cur, []int{x, y})
		}
	}
}

func TestSample1(t *testing.T) {
	s := `8 10 8
1 1 10
1 4 13
1 7 1
1 8 2
2 2 0
2 5 14
2 6 0
2 6 1
4 8
10 5
8 8
10 6
10 2
1 8
7 8
10 6
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 2 3
1 1 2
2 1 1
1 1 5
1 3
2 1
1 3
`
	runSample(t, s)
}
