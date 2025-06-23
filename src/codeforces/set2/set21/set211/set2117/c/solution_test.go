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
	runSample(t, `6
1 2 2 3 1 5
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `8
1 2 1 3 2 1 3 2
3`)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
5 4 3 2 1
1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `10
5 8 7 5 8 5 7 8 10 9
3`)
}
