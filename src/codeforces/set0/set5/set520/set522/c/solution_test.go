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
	s := `3 4
2 3 2 1
1 0
0 0`
	expect := "YNNY"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 5
1 2 1 3 1
3 0
0 0
2 1
4 0`
	expect := "YYYNY"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 2
2 2
0 0
0 0
1 1
`
	expect := "NY"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `3 5
2 2 3 4 1
0 0
0 0
`
	expect := "YYNNY"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `3 3
1 1 1
0 0
1 1
`
	expect := "YYY"
	runSample(t, s, expect)
}
