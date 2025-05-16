package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	ans := readNum(reader)
	if res != ans {
		t.Errorf("Sample expect %d, but got %d", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 7 6 11
3
5 3 8
6 7 11
5 2 5
4`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 4 3 10
3
3 1 4
4 5 9
3 10 10
6`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `1 1 2 10
2
1 1 3
2 6 10
-1`
	runSample(t, s)
}

