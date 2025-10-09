package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := drive(reader)

	for i, x := range ans {
		var y int
		fmt.Fscan(reader, &y)
		if x != y {
			t.Fatalf("Sample expect %d, but got %d at %d", y, x, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 6
1 2 3 4 6
1
2
3
4
5
1
0
1
3
5
6
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10 9
151790 360570 1 1 123690 162690 1 155208 227488 1
3
10
10
7
9
10
3
4
9
0
1
0
1
3
6
3
6
3`
	runSample(t, s)
}
