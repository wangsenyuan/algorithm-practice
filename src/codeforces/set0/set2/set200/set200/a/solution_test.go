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

	for i, cur := range res {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		if cur[0] != x || cur[1] != y {
			t.Fatalf("Sample expect (%d %d), but got %v for %d-th customer", x, y, cur, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 6
1 1
1 1
1 1
1 2
1 3
1 3
1 1
1 2
2 1
1 3
1 4
2 3
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 3 12
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
2 2
1 2
2 1
2 3
3 2
1 1
1 3
3 1
3 3
4 2
4 1
4 3
`
	runSample(t, s)
}
