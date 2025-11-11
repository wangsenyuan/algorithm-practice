package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	ws, res := drive(reader)
	if expect != (res != nil) {
		t.Fatalf("Sample expect %t, but got %v", expect, res)
	}
	if !expect {
		return
	}
	type pair struct {
		first  int
		second int
	}
	id := make(map[pair]int)

	for i, cur := range res {
		x, y := cur[0], cur[1]
		if y-x != ws[i] {
			t.Fatalf("Sample result failed at %d, expect %d, but got %d", i, ws[i], y-x)
		}
		if x > 0 {
			left := pair{x - 1, y}
			if _, ok := id[left]; !ok {
				t.Fatalf("Sample result failed at %d, (%d %d) should be before (%d %d)", i, x-1, y, x, y)
			}
		}
		if y > 0 {
			down := pair{x, y - 1}
			if _, ok := id[down]; !ok {
				t.Fatalf("Sample result failed at %d, (%d %d) should be before (%d %d)", i, x, y-1, x, y)
			}
		}
		id[pair{x, y}] = i
	}
}

func TestSample1(t *testing.T) {
	s := `5
2 0
0 0
1 0
1 1
0 1
0 -1 -2 1 0
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
1 0
0 0
2 0
0 1 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `31
0 0
0 1
0 2
0 3
1 0
1 1
2 0
2 1
3 0
4 0
5 0
6 0
7 0
8 0
9 0
10 0
11 0
12 0
13 0
14 0
15 0
16 0
17 0
18 0
19 0
20 0
21 0
22 0
23 0
24 0
25 0
0 -1 -2 -3 -4 -5 -6 -7 -8 -9 -10 -11 -12 -13 -14 1 -15 2 -16 -17 -18 3 -19 -20 0 -21 -22 -23 -24 -25 -1
`
	expect := true
	runSample(t, s, expect)
}
