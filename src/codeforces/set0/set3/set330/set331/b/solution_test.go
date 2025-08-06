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
	runSample(t, `5
1 3 4 2 5
6
1 1 5
1 3 4
2 2 3
1 1 5
2 1 5
1 1 5
2
1
3
5`)
}
