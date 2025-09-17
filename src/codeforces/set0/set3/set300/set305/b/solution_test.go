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
	runSample(t, `9 4
2
2 4
`, true)
}

func TestSample2(t *testing.T) {
	runSample(t, `9 4
3
2 3 1
`, true)
}

func TestSample3(t *testing.T) {
	runSample(t, `9 4
3
1 2 4
`, false)
}

func TestSample4(t *testing.T) {
	runSample(t, `2 1
4
2 1 1 1
`, false)
}

func TestSample5(t *testing.T) {
	runSample(t, `913255926290448385 4400000000
2
4400000000 4400000000
`, false)
}
