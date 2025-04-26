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
			t.Fatalf("Sample expect %v, but got %v", expect, ans)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `6 5
1 2
5 4
2 3
3 1
3 6
6
1 3
2 5
1 5
5 5
2 4
3 3
`
	expect := `4
5
6
3
4
2
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 10
8 6
8 7
8 3
3 7
4 8
1 6
5 1
8 7
6 8
1 6
13
1 10
2 6
3 3
5 5
2 2
1 3
10 10
7 7
2 4
3 6
2 7
9 9
3 6
`
	expect := `8
4
2
3
2
2
2
3
3
4
5
2
4
`
	runSample(t, s, expect)
}
