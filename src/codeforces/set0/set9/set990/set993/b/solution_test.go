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
	runSample(t, `2 2
1 2 3 4
1 5 3 4
	`, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, `2 2
1 2 3 4
1 5 6 4
	`, 0)
}

func TestSample3(t *testing.T) {
	runSample(t, `2 3
1 2 4 5
1 2 1 3 2 3
	`, -1)
}

func TestSample4(t *testing.T) {
	runSample(t, `4 4
1 2 3 4 5 6 7 8
2 3 4 5 6 7 8 1
	`, -1)
}

func TestSample5(t *testing.T) {
	runSample(t, `9 1
3 4 3 2 3 7 3 5 9 4 1 9 6 4 5 2 7 6
8 3
	`, 3)
}
