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
	runSample(t, `1 1
1
1`)
}
func TestSample2(t *testing.T) {
	runSample(t, `2 2
10
11
2`)
}

func TestSample3(t *testing.T) {
	runSample(t, `4 3
100
011
000
101
2`)
}

func TestSample4(t *testing.T) {
	runSample(t, `11 16
0111110101100011
1000101100010000
0010110110010101
0110110010110010
0011101101110000
1001100011010111
0010011111111000
0100100100111110
1001000000100111
0110000011001000
1011111011010000
9`)
}
