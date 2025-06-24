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
	runSample(t, `3
7 2
1 6
5`)
}

func TestSample2(t *testing.T) {
	runSample(t, `1
41 42
13 56
42`)
}

func TestSample3(t *testing.T) {
	runSample(t, `100
100 99
199 1
0`)
}

func TestSample4(t *testing.T) {
	runSample(t, `96929423
5105216413055191 10822465733465225
1543712011036057 14412421458305526
79154049`)
}
