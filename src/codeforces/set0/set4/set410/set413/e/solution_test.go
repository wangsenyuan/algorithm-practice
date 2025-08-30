package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	for _, y := range res {
		x := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %s, but got %v", s, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `4 7
.X..
...X
5 1
1 3
7 7
1 4
6 1
4 7
5 7
1
4
0
5
2
2
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `10 3
X...X..X..
..X...X..X
11 7
7 18
18 10
9
-1
3`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1 1
.
.
1 2
1`
	runSample(t, s)
}
