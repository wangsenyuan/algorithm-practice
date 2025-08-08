package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	for _, x := range ans {
		var y int
		fmt.Fscan(reader, &y)
		if x != y {
			t.Fatalf("Sample expect %d, but got %d", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 3
1 2 1
1 3 3
1 4 4
3 1 4 0
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `8 4
3 5 4
3 7 6
2 8 8
1 8 1
0 8 4 6 4 8 6 1
	`)
}
