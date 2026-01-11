package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `5
0 0
0 2
2 0
2 2
1 1
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `9
0 0
1 1
2 2
0 1
1 0
0 2
2 0
1 2
2 1
`, 5)
}

func TestSample3(t *testing.T) {
	runSample(t, `54
0 8
3 2
9 3
7 2
8 2
2 8
10 10
7 6
1 1
9 7
4 0
6 10
10 1
10 8
5 1
0 4
7 10
3 6
0 5
4 3
3 0
5 10
6 9
5 4
6 6
8 5
0 7
5 8
1 2
2 2
9 4
2 4
0 10
5 9
10 9
7 9
9 9
2 5
4 10
8 9
7 7
5 2
6 5
4 1
10 6
6 3
9 6
0 9
7 3
7 5
8 4
1 3
0 3
2 10
`, 14)
}
