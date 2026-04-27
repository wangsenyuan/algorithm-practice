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
	s := `2 2
1 2`
	expect := 2
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 3
1 2
1 3
1 4
1 5`
	expect := 9
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `6 3
1 2
1 3
2 4
2 5
3 6`
	expect := 17
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `10 5
5 6
4 9
3 9
2 6
2 8
8 9
6 10
1 6
4 7`
	expect := 35
	runSample(t, s, expect)
}
