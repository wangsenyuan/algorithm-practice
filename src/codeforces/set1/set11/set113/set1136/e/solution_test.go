package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	reader = bufio.NewReader(strings.NewReader(expect))
	for _, x := range ans {
		y := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %v, but got %v", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 2 3
1 -1
5
s 2 3
+ 1 2
s 1 2
+ 3 1
s 2 3
`
	expect := `5
7
8`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `3
3 6 7
3 1
3
+ 1 3
+ 2 4
s 1 3
`
	expect := `33`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10
-100 -88 -84 -47 -31 -4 13 18 18 10
6 -3 5 -6 -4 -8 -2 -6 -10
2
s 1 4
s 1 6
`
	expect := `-319
-354`
	runSample(t, s, expect)
}
