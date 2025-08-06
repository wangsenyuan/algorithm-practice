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
		if y != x {
			t.Errorf("Sample expect %v, but got %v", y, x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 10
1 2 7
3 6 5
3 4 8
7
11
8`)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 10
1 2 5
7 4 5
5
9`)
}

func TestSample3(t *testing.T) {
	runSample(t, `10 10
7 5 9
10 9 8
7 5 10
4 2 8
2 9 1
2 8 10
7 10 9
7 7 2
5 1 5
4 7 9
14
15
29
27
38
29
30
21
42
14`)
}
