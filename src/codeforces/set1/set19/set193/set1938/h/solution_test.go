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
	runSample(t, `4
11101101
00
10001
10
`, 5)
}

func TestSample2(t *testing.T) {
	runSample(t, `2
101010
010101
`, 6)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
0000
11
0
00000000
1
`, 0)
}

func TestSample4(t *testing.T) {
	runSample(t, `2
000
001000
`, 4)
}

func TestMixedTablesCanChooseBothColorsSeveralTimes(t *testing.T) {
	runSample(t, `4
00001
00001
01111
01111
`, 4)
}
