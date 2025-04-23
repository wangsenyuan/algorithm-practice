package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	ans := process(reader)
	if ans != expect {
		t.Errorf("Sample expect %d, but got %d", expect, ans)
	}
}

func TestSample1(t *testing.T) {
	s := `3 3 10
1 9 4
2 3 0
1 2
3 2
1 3`
	expect := 6
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `2 1 10
1 2
4 6
1 2`
	expect := -1
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `8 7 25
22 14 5 3 10 14 11 1
9 5 4 10 7 16 18 18
2 8
6 3
3 5
7 5
2 6
1 4
4 7`
	expect := 52
	runSample(t, s, expect)
}
