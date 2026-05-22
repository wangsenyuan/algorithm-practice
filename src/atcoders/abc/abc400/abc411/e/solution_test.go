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
		t.Errorf("Sample result %d, expect %d", res, expect)
	}
}

func TestSample1(t *testing.T) {
	s := `2
1 1 4 4 4 4
1 1 1 3 3 3
`
	expect := 332748121
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8
55 76 80 21 34 28
82 84 2 32 56 17
11 57 37 28 39 18
47 2 97 25 75 29
72 45 22 75 26 81
6 79 16 68 68 40
31 80 68 57 18 55
49 10 63 91 93 40
`
	expect := 213725517
	runSample(t, s, expect)
}
