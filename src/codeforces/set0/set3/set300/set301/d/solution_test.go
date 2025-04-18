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
	s := `1 1
1
1 1
`
	expect := "1"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 9
1 2 3 4 5 6 7 8 9 10
1 10
2 9
3 8
4 7
5 6
2 2
9 10
5 10
4 10
`
	expect := `27
14
8
4
2
1
2
7
9
`
	runSample(t, s, expect)
}
