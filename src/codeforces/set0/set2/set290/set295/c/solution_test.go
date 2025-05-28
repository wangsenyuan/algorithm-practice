package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	expect_min := readNum(reader)
	expect_ways := readNum(reader)
	if res[0] != expect_min || res[1] != expect_ways {
		t.Errorf("Sample expect %d, %d, but got %d, %d", expect_min, expect_ways, res[0], res[1])
	}
}

func TestSample1(t *testing.T) {
	s := `1 50
50
1
1`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 100
50 50 100
5
2`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `2 50
50 50
-1
0`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `5 258
100 100 50 50 50
3
72`
	runSample(t, s)
}
