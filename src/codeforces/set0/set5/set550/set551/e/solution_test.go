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
	runSample(t, `4 3
1 2 3 4
1 1 2 1
1 1 1 1
2 3
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 3
1 2
1 2 2 1
2 3
2 4
0
-1`)
}