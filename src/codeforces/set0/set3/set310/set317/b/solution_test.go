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
	s := `1 3
0 1
0 0
0 -1
0
1
0`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6 5
0 -2
0 -1
0 0
0 1
0 2
0
1
2
1
0`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `30000 19881
-70 -70
3`
	runSample(t, s)
}
