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
		if x != cur[0] || y != cur[1] {
			t.Errorf("Sample expect %v, but got %v", cur, []int{x, y})
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3 3 1 1 9
1 1
1 2
1 3
2 1
2 2
2 3
3 1
3 2
3 3
1 3
1 2
1 1
2 3
2 2
2 1
3 3
3 2
3 1
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5 5 0 0 0 1
1 4
1 4
`
	runSample(t, s)
}
