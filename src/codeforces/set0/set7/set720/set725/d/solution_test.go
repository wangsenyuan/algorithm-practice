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
	s := `8
20 1000
32 37
40 1000
45 50
16 16
16 16
14 1000
2 1000
3
`
	runSample(t, s)
}

func TestSample2(t *testing.T) {
	s := `7
4 4
4 4
4 4
4 4
4 4
4 4
5 5
2`
	runSample(t, s)
}

func TestSample3(t *testing.T) {
	s := `7
14000000003 1000000000000000000
81000000000 88000000000
5000000000 7000000000
15000000000 39000000000
46000000000 51000000000
0 1000000000
0 0
2`
	runSample(t, s)
}

func TestSample4(t *testing.T) {
	s := `9
4 70
32 56
32 65
77 78
5 29
72 100
0 55
42 52
66 72
7`
	runSample(t, s)
}

func TestSample5(t *testing.T) {
	s := `2
100 150
5 100000
1`
	runSample(t, s)
}
