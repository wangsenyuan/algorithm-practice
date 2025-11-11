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
	s := `6 2
38 56 49
7 3 4
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `8 2
1 22 3 44
5 4 3 2
`
// 0...99 100
// 50....59 = 10
// 0 22 44 66 88 = 5
// 0 3..... 99 = 100 / 3 = 34
// 30 33 36 39 = 4
// 0 44, 88 = 3

	expect := 32400
	runSample(t, s, expect)
}
