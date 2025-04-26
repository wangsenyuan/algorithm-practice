package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res := process(bufio.NewReader(strings.NewReader(s)))

	reader := bufio.NewReader(strings.NewReader(expect))
	for _, x := range res {
		y := readNum(reader)
		if x != y {
			t.Fatalf("Sample expect %v, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `13
3 1
3 2
3 3
1
3 4
2
3 5
1
3 6
2
3 7
2
1`
	expect := `1
5
14
11
27
23
48
38
74
73
122
102
88
`
	runSample(t, s, expect)
}
