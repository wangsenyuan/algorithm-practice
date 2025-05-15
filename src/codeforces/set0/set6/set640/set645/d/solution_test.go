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
		t.Fatalf("Sample expect %d, but got %d", ans, res)
	}
}

func TestSample1(t *testing.T) {
	s := `4 5
2 1
1 3
2 3
4 2
4 3
4
	`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `3 2
1 2
3 2
-1
	`
	runSample(t, s)
}
