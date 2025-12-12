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
	s := `15
0 1 0 0 1 1 0 1 1 1 1 1 0 1 0
10 7
10 3
10 8
5 7
13 14
8 13
15 4
15 13
5 2
9 3
11 15
13 6
1 12
9 1
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `11
0 0 0 1 1 0 1 0 0 1 1
1 2
1 3
2 4
2 5
5 6
5 7
3 8
3 9
3 10
9 11
`
	expect := 2
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4
0 0 0 0
1 2
2 3
3 4
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `42
1 0 0 1 0 1 1 0 1 1 1 1 0 1 0 0 0 1 1 1 0 1 1 1 0 1 0 0 0 0 0 0 1 0 0 0 0 0 0 1 0 0
35 6
35 39
4 31
31 5
14 35
1 2
40 32
35 31
37 35
32 38
1 36
3 25
35 11
26 35
24 35
3 2
35 23
21 1
20 27
16 26
2 18
34 39
39 28
3 32
26 30
41 7
13 35
1 8
31 22
33 21
21 29
28 10
2 19
2 17
27 24
9 1
42 1
1 15
1 35
12 2
41 1
`
	expect := 3
	runSample(t, s, expect)
}
