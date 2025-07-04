package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect := readNum(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4
1 4 2
4
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4 6
2 -1 2 -1
1
`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `8 10
1 -2 3 -4 5 -6 7 -8
2
`
	runSample(t, s)
}

// func TestSample4(t *testing.T) {
// 	s := `100 4064
// -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100 -100 100
// -100
// `
// 	runSample(t, s)
// }
