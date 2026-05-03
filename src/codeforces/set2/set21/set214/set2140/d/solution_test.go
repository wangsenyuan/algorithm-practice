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
	runSample(t, `2
1 1000000000
1 1000000000`, 2999999997)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
1 10
2 15
3 9`, 42)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
1 11
2 7
15 20
1 3
11 15`, 59)
}

func TestSample4(t *testing.T) {
	runSample(t, `1
1000000000 1000000000`, 0)
}

func TestSample5(t *testing.T) {
	// 13 + 7 = 20
	runSample(t, `3
1 4
1 8
4 7`, 20)
}
