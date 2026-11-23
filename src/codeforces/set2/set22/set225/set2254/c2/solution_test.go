package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect int) {
	t.Helper()
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4
0100
0001
`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `4
0100
0010
`, -1)
}

func TestSample3(t *testing.T) {
	runSample(t, `6
110000
000011
`, 4)
}

func TestSample4(t *testing.T) {
	runSample(t, `8
10101010
10101010
`, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, `5
01001
10010
`, 3)
}

func TestSample6(t *testing.T) {
	runSample(t, `7
1110001
1110100
`, 1)
}

func TestSample7(t *testing.T) {
	runSample(t, `7
1001100
1011001
`, -1)
}
