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
	runSample(t, `5
1 2 3 6 8
2
`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3
3 6 5
-1
`)
}
func TestSample3(t *testing.T) {
	runSample(t, `6
2 4 8 6 10 18
3
`)
}
