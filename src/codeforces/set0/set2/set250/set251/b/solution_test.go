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
	s := `4 1
2 3 4 1
1 2 3 4
`
	expect := false
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 1
4 3 1 2
3 4 2 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `4 3
4 3 1 2
3 4 2 1
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 2
4 3 1 2
2 1 4 3
`
	expect := true
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `5 3
2 1 4 3 5
2 1 4 3 5
`
	expect := false
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := `8 9
2 3 1 5 6 7 8 4
2 3 1 4 5 6 7 8
`
	expect := true
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := `55 28
34 11 18 6 16 43 12 25 48 27 35 17 19 14 33 30 7 53 52 2 15 10 44 1 37 28 22 49 46 8 45 39 21 47 40 20 41 51 13 24 42 55 23 4 36 38 50 31 3 9 54 32 5 29 26
34 11 18 6 16 43 12 25 48 27 35 17 19 14 33 30 7 53 52 2 15 10 44 1 37 28 22 49 46 8 45 39 21 47 40 20 41 51 13 24 42 55 23 4 36 38 50 31 3 9 54 32 5 29 26
`
	expect := true
	runSample(t, s, expect)
}
