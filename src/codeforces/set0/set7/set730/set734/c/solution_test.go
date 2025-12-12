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
	s := `20 3 2
10 99
2 4 3
20 10 40
4 15
10 80`
	expect := 20
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `20 3 2
10 99
2 4 3
200 100 400
4 15
100 800
`
	expect := 200
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `10 3 3
10 33
1 7 6
17 25 68
2 9 10
78 89 125
`
	expect := 10
	runSample(t, s, expect)
}
