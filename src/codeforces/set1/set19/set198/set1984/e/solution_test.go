package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))

	res := process(reader)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5
1 2
1 3
2 4
2 5`
	expect := 4
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5
1 2
2 3
3 4
4 5`
	expect := 3
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6
1 2
1 3
1 4
1 5
1 6`
	expect := 5
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10
9 3
8 1
10 6
8 5
7 8
4 6
1 3
10 1
2 7`
	expect := 6
	runSample(t, s, expect)
}
