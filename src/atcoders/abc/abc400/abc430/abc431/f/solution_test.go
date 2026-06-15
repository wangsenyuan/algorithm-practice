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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 1
5 2 1 2
`
	expect := 3
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 10
20 40 60 80 100
`
	expect := 1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `15 12345
18270 31252 27543 31406 22271 13402 12279 25697 18349 27615 39360 22790 32581 23990 36154
`
	expect := 858152905
	runSample(t, s, expect)
}
