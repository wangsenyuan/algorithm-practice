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
		if y != x {
			t.Errorf("Sample expect %v, but got %v", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `5 10
1 2 0 -1 3
3 1 5 -2 0
2 5
1 3 3 3
2 5
2 4
2 1
1 2 1 4
2 1
2 4
1 4 2 1
2 2
`
	expect := `0
3
-1
3
2
3
-1
`
	runSample(t, s, expect)
}
