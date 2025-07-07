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
	runSample(t, `2 3
ATK 2000
DEF 1700
2500
2500
2500
3000`)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 4
ATK 10
ATK 100
ATK 1000
1
11
101
1001
992`)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 4
DEF 0
ATK 0
0
0
1
1
1`)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 6
DEF 0
DEF 0
DEF 0
DEF 0
DEF 0
1
1
1
1
1
1
1`)
}
