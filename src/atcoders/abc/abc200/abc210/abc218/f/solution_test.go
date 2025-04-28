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
			t.Fatalf("Sample expect %s, but got %v", expect, res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
1 2
1 3
2 3
`
	expect := `1
2
1
`
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 4
1 2
2 3
2 4
3 4
`
	expect := `-1
2
3
2
`
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4
1 2
2 3
2 4
3 4
`
	expect := `-1
2
3
2
`
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `5 10
1 2
1 4
1 5
2 1
2 3
3 1
3 2
3 5
4 2
4 3
`
	expect := `1
1
3
1
1
1
1
1
1
1
`
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `4 1
1 2
`
	expect := `-1`
	runSample(t, s, expect)
}
