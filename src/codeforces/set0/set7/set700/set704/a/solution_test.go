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
			t.Fatalf("Sample expect %s, but got %v", s, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
1 3
1 1
1 2
2 3
1
2
3
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 6
1 2
1 4
1 2
3 3
1 3
1 3
1
2
3
0
1
2
`
	runSample(t, s)
}
