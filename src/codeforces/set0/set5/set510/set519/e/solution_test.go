package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
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
1 2
1 3
2 4
1
2 3
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4
1 2
2 3
2 4
2
1 2
1 3
0
2`
	runSample(t, s)
}
