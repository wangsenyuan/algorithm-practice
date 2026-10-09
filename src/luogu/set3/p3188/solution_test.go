package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res, ok := drive(bufio.NewReader(strings.NewReader(s)))
	if !ok || res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 10
8 9
5 8
4 6
2 5
`, 14)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 13
8 9
5 8
4 6
2 5
`, 19)
}

func TestSample3(t *testing.T) {
	runSample(t, `16 75594681
393216 5533
2 77
32768 467
29360128 407840
112 68
24576 372
768 60
33554432 466099
16384 318
33554432 466090
2048 111
24576 350
9216 216
12582912 174768
16384 295
1024 76
`, 1050650)
}
