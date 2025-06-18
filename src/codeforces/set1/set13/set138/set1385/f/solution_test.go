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
	runSample(t, `8 3
1 2
1 5
7 6
6 8
3 1
6 4
6 1
2
	`)
}

func TestSample2(t *testing.T) {
	runSample(t, `10 3
1 2
1 10
2 3
1 5
1 6
2 4
7 10
10 9
8 10
3
	`)
}

func TestSample3(t *testing.T) {
	runSample(t, `7 2
3 1
4 5
3 6
7 4
1 2
1 4
3
	`)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 1
1 2
2 3
4 3
5 3
4
	`)
}
