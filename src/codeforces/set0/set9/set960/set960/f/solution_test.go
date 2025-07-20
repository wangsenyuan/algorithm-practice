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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
3 1 3
1 2 1
2 3 2
2`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `5 5
1 3 2
3 2 3
3 4 5
5 4 0
4 5 8
3`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `5 10
3 4 8366
5 1 6059
2 1 72369
2 2 35472
5 3 50268
2 4 98054
5 1 26220
2 3 24841
1 3 42450
3 1 59590
3`
	runSample(t, s)
}
