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
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6
1 2
2 3
2 4
4 5
1 6
3`)
}

func TestSample2(t *testing.T) {
	runSample(t, `7
1 2
1 3
3 4
1 5
5 6
6 7
-1`)
}

func TestSample3(t *testing.T) {
	runSample(t, `3
3 1
1 2
1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `11
8 9
2 7
1 11
3 2
9 1
8 5
8 6
5 4
4 10
8 3
1`)
}
