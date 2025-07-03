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
	runSample(t, `9
8 6 4 2 1 4 7 10 2
3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `9
-1 6 -1 2 -1 4 7 -1 2
3`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
-1 -1 -1 -1 -1
1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `7
-1 -1 4 5 1 2 3
2`)
}

func TestSample5(t *testing.T) {
	runSample(t, `6
-1 6 1 -1 -1 -1
2`)
}

func TestSample6(t *testing.T) {
	runSample(t, `7
-1 2 4 -1 4 1 5
3`)
}
