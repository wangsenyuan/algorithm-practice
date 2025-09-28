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
	s := `3 5 1 3
1 1 2
2 2 3
3 3 3
4 1 1
10 1 3
`
	expect := "XXRR"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 3 2 1
1 1 2
2 1 2
4 1 2
`
	expect := "XXL"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 4 3 4
1 2 4
2 1 2
3 3 4
4 2 3
`
	expect := "XR"
	runSample(t, s, expect)
}
