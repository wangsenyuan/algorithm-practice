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
	s := `4
2 1 6 4
3 4 4 2
3`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `4
10 5 6 4
1 11 4 2
6`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `3
10 1 10
1 10 1 1
3`
	// 从这个例子里面，可以到到还是有点技巧的
	runSample(t, s)
}
