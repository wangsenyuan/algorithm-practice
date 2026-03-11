package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 4
WWBW
BWWW
WWWB
`, false)
}

func TestSample2(t *testing.T) {
	runSample(t, `3 1
B
B
W
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 5
WBBBW
WBBBW
WBBWW
WBBBW
WWWWW`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 5
WBBBW
WBBWW
WBBWW
BBBWW
BBWWW`, false)
}
