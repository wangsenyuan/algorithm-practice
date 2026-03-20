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
	for i, x := range res {
		var y int
		fmt.Fscan(reader, &y)
		if x != y {
			t.Fatalf("Sample expect %d, but got %d at %d", y, x, i)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6
pasha 0
gerald 1
gerald 1
valera 2
igor 3
olesya 1
5
1 1
1 2
1 3
3 1
6 1
2
2
0
1
0
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `6
valera 0
valera 1
valera 1
gerald 0
valera 4
kolya 4
7
1 1
1 2
2 1
2 2
4 1
5 1
6 1
1
0
0
0
2
0
0
`
	runSample(t, s)
}
