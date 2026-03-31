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
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `4 6
1 0 0 6`, 2)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 2
2 1`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `5 24
0 0 4 0 0`, 10)
}

func TestSample4(t *testing.T) {
	runSample(t, `5 6
0 0 6 0 0`, 0)
}

func TestSample5(t *testing.T) {
	runSample(t, `20 2000
1 0 0 0 0 14 0 0 0 0 0 0 0 0 0 514 0 0 0 0`, 973702700)
}
