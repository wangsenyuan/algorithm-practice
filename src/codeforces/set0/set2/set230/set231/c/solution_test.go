package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	ans := readNNums(reader, 2)
	if res[0] != ans[0] || res[1] != ans[1] {
		t.Errorf("Sample expect %v, but got %v", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `5 3
6 3 4 0 2
3 4`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 4
5 5 5
3 5`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `5 3
3 1 2 2 1
4 2`
	runSample(t, s)
}
