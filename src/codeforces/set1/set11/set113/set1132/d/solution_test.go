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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `2 4
3 2
4 2
5`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `1 6
4
2
2`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 2
2 10
3 15
-1`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `1 1
1
2
0`
	runSample(t, s)
}
