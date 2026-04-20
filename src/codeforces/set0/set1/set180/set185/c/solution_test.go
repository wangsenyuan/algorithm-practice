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
	s := `1
1
2
`
	expect := "Fat Rat"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2
2 2
1 2
4
`
	expect := "Cerealguy"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `2
2 2
1 2
5
`
	expect := "Fat Rat"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `6
1 1 2 2 1 1
1 1 1 2 1 1
1 1 2 4 2
1 9 2 2
1 9 4
4 4
8
`
	expect := "Cerealguy"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `6
1 1 2 2 1 1
1 1 2 2 1 1
2 4 2 4 2
2 2 2 2
4 10 4
4 4
8
`
	expect := "Fat Rat"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `3
3 2 2
3 5 1
3 1
5
`
	expect := "Cerealguy"
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `6
1 1 1 1 1 1
1 2 1 1 2 1
1 9 1 9 1
1 1 1 1
2 9 2
2 2
4
`
	expect := "Fat Rat"
	runSample(t, s, expect)
}
