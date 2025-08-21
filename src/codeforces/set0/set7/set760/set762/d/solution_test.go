package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)

	if res != expect {
		t.Fatalf("sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3
1 1 1
1 -1 1
1 1 1
`
	runSample(t, s, 7)
}

func TestSample2(t *testing.T) {
	s := `5
10 10 10 -1 -1
-1 10 10 10 10
-1 10 10 10 10
`
	runSample(t, s, 110)
}

func TestSample3(t *testing.T) {
	s := `20
16 82 25 21 -60 9 29 -55 70 54 -50 10 -19 40 46 41 31 -66 1 85
-15 75 -94 -7 -50 -97 -55 -24 44 -69 -73 15 -9 98 92 -92 72 -32 -46 59
74 99 -6 97 -59 41 -22 -8 -27 75 3 -56 -38 -56 -43 16 -43 -92 55 -63
`
	runSample(t, s, 946)
}
