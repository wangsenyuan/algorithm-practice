package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4
3 2 15 9
`
	expect := 27
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4
8 2 2 7
`
	expect := 18
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `100
17 23 71 25 50 71 85 46 78 72 89 26 23 70 40 59 23 43 86 81 70 89 92 98 85 88 16 10 26 91 61 58 23 13 75 39 48 15 73 79 59 29 48 32 45 44 25 37 58 54 45 67 27 77 20 64 95 41 80 53 69 24 38 97 59 94 50 88 92 47 95 31 66 48 48 56 37 76 42 74 55 34 43 79 65 82 70 52 48 56 36 17 14 65 77 81 88 18 33 40
`
	expect := 5030
	runSample(t, s, expect)
}

