package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2
2 3
3 4
4 5
`
	expect := "6"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
1 2
1 3
2 4
2 5
3 6
3 7
6 8
`
	expect := "18"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3
1 2
1 3
`
	expect := "3"
	runSample(t, s, expect)
}

func TestSingleNode(t *testing.T) {
	s := `1
`
	expect := "1"
	runSample(t, s, expect)
}

func TestStarTree(t *testing.T) {
	s := `5
1 2
1 3
1 4
1 5
`
	expect := "5"
	runSample(t, s, expect)
}
