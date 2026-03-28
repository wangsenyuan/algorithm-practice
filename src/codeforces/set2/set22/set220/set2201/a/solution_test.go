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
	runSample(t, `5
1 2 3 4 5
`, 15)
}

func TestSample2(t *testing.T) {
	runSample(t, `5
1 3 5 7 9
`, 35)
}

func TestSample3(t *testing.T) {
	runSample(t, `5
1 2 5 6 5
`, 25)
}

func TestSample4(t *testing.T) {
	runSample(t, `7
1 2 4 5 3 7 8
`, 60)
}

func TestSample5(t *testing.T) {
	runSample(t, `9
9 8 9 2 3 4 4 5 3
`, 78)
}
