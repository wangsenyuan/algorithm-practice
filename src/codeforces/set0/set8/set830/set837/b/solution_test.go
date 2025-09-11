package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	res := drive(bufio.NewReader(strings.NewReader(s)))
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `6 5
RRRRR
RRRRR
BBBBB
BBBBB
GGGGG
GGGGG
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `4 3
BRG
BRG
BRG
BRG
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `6 7
RRRGGGG
RRRGGGG
RRRGGGG
RRRBBBB
RRRBBBB
RRRBBBB
`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 4
RRRR
RRRR
BBBB
GGGG
`, false)
}
