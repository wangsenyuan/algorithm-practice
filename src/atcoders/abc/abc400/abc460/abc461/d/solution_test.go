package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 4 3
1001
1101
0110
`
	expect := 8
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `5 4 20
0101
1010
0101
1010
0101
`
	expect := 0
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `15 20 17
10111101101100000100
01100000000010000011
01110010111000111000
11001100000111011000
10100001100011100010
01101000101010000101
10110001111110000100
10110011101100101101
01010001110011001001
01010110010001100110
01110100011110011110
01100000100111010010
11001101100111101100
10111100010101111011
00101101011100010000
`
	expect := 448
	runSample(t, s, expect)
}
