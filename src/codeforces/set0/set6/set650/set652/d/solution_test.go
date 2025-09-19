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
	s := `4
1 8
2 3
4 7
5 6
3
0
1
0
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3
3 4
1 5
2 6
0
1
1
`
	runSample(t, s)
}
