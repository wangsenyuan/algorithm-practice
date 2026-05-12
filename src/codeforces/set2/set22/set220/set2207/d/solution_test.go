package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `6 2 1
1 2
2 3
2 4
1 5
5 6
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `7 1 4
1 2
2 3
3 4
4 5
5 6
6 7
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 1 3
1 3
2 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 1 4
1 3
3 4
4 2
`
	expect := false
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `9 3 5
4 5
5 6
4 7
9 8
8 7
1 2
2 3
3 4
`
	expect := false
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `9 4 5
4 5
5 6
4 7
9 8
8 7
1 2
2 3
3 4
`
	expect := true
	runSample(t, s, expect)
}